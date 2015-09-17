package main

import (
    "net/http"

    "github.com/gorilla/mux"

    "github.com/javinc/playgo/goryo/services/test"
    "github.com/javinc/playgo/goryo/services/user"

    "github.com/javinc/playgo/goryo/controllers"
)

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", controllers.IndexHandler)
    r.HandleFunc("/test", controllers.TestHandler)
    r.HandleFunc("/user", user.Handler)
    r.HandleFunc("/login", test.Login)

    // Bind to a port and pass our router in
    http.ListenAndServe(":8000", r)
}
