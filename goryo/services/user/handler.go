package user

import (
    "net/http"

    "github.com/javinc/playgo/goryo/resources/user"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    output := "user services \n"

    // initialize object
    m := user.Model{}

    // lets base on the request type
    switch r.Method {
    case "GET":
        Find(m, w, r)
        return
    case "POST":
        Create(m, w, r)
        return
    }

    // catcher
    output += r.Method + " method done! \n"

    w.Write([]byte(output))
}
