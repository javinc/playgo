package services

import (
    "net/http"

    "github.com/javinc/playgo/rango/resources"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
    output := "test services \n"

    // initialize object
    t := resources.Test{}

    // excute mehtod of object
    output += t.Find("something awesome!")
    // output += t.Create()
    output += "\n"

    w.Write([]byte(output))
}
