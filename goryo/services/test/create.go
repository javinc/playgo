package test

import (
    "net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
    output := "create test services \n"

    // excute mehtod of object
    output += TestResource.Create("something awesome!")
    output += "\n"

    w.Write([]byte(output))
}
