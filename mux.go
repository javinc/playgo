package main

import (
    "github.com/javinc/acube/api/helper"
    "net/http"
    "net/url"
    "io/ioutil"
)

func muxHandler() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        method := r.Method;

        render := "URL is " + r.URL.String() + "\n"
        render += "Content-Type is " + r.Header["Content-Type"][0] + "\n"
        render += "Method is " + method + "\n"

        // get GET params
        render += "\nGET params: \n"

        getParam := r.URL.Query()
        for key, value := range getParam {
            render += "\t" + key + ":" + value[0] + "\n"
        }

        // get POST params
        render += "\nPOST params: \n"

        body, _ := ioutil.ReadAll(r.Body)
        postParam, _ := url.ParseQuery(string(body))
        for key, value := range postParam {
            render += "\t" + key + ":" + value[0] + "\n"
        }

        // render to page
        w.Write([]byte(render))
    })
}

func main() {
    mux := http.NewServeMux()

    handler := muxHandler()
    mux.Handle("/", handler)

    helper.ConsoleLog("Listening...")
    http.ListenAndServe(":3000", mux)
}
