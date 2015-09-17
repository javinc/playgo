package test

import (
    "net/http"

    "github.com/javinc/playgo/limbo/resources/test"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    output := "test services \n"

    // initialize object
    t := test.Model{}

    // lets base on the request type
    switch r.Method {
    case "GET":
        Find(t, w, r)
        return
    case "POST":
        Create(t, w, r)
        return
    }

    // catcher
    output += r.Method + " method done! \n"

    w.Write([]byte(output))
}
