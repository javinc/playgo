package main

import (
    "net/http"
    "net/url"
    "io/ioutil"

    "github.com/gorilla/mux"
    "github.com/javinc/playgo/rango"
)

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", IndexHandler)
    r.HandleFunc("/some", SomeHandler)

    // Bind to a port and pass our router in
    http.ListenAndServe(":8000", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(r.Method + " Gorilla!\n"))
}

func SomeHandler(w http.ResponseWriter, r *http.Request)  {
    render := "URL is " + r.URL.String() + "\n"
    render += "Method is " + r.Method + "\n"

    getParam := r.URL.Query()
    // get GET params
    render += "\nGET params: \n"
    for key, value := range getParam {
        render += "\t" + key + ":" + value[0] + "\n"
    }

    render += "\nPOST params: \n"
    // get POST params
    body, _ := ioutil.ReadAll(r.Body)
    postParam, _ := url.ParseQuery(string(body))
    for key, value := range postParam {
        render += "\t" + key + ":" + value[0] + "\n"
    }

    w.Write([]byte(render))
}
