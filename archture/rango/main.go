package main

import (
    "net/http"
    "net/url"
    "io/ioutil"

    "github.com/gorilla/mux"
    "github.com/javinc/playgo/rango/services"
)

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", IndexHandler)
    r.HandleFunc("/user", services.UserHandler)
    r.HandleFunc("/test", services.TestHandler)

    // Bind to a port and pass our router in
    http.ListenAndServe(":8000", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    render := r.Method + " Gorilla!\n"
    render += "URL is " + r.URL.String() + "\n"

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
