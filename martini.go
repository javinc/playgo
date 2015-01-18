package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/hello", func(req *http.Request) string {

		var response string

		response = "GET\n\n"

		v := req.URL.Query()

		for key, value := range v {
			fmt.Printf("key: %v\n", key)
			fmt.Printf("value: %v\n", value)

			if len(value) > 1 {
				for _, arrayValue := range value {
					response += key + ":" + arrayValue + "\n"
				}
			} else {
				response += key + ":" + v.Get(key) + "\n"
			}
		}

		return response
	})

	m.Post("/hello-post", func(req *http.Request) string {

		var response string

		response = "POST\n\n"

		// if we would use this, then the POST and GET requests would be merged
		// to the req.Form variable
		// req.ParseForm()
		// v := req.Form

		body, _ := ioutil.ReadAll(req.Body)
		v, _ := url.ParseQuery(string(body))

		for key, value := range v {
			fmt.Printf("key: %v\n", key)
			fmt.Printf("value: %v\n", value)

			if len(value) > 1 {
				for _, arrayValue := range value {
					response += key + ":" + arrayValue + "\n"
				}
			} else {
				response += key + ":" + v.Get(key) + "\n"
			}
		}

		return response
	})

	m.Run()
}
