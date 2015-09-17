package test

import (
    "net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
    output := "login test services \n"

    output += "you are logging in! \n"

    w.Write([]byte(output))
}
