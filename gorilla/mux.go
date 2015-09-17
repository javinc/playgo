package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", IndexHandler)

    // Bind to a port and pass our router in
    http.ListenAndServe(":8000", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(r.Method + " Gorilla!\n"))
}
