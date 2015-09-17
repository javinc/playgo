package test

import (
    "net/http"

    "github.com/javinc/playgo/goryo/resources/test"
)

func Create(m test.Model, w http.ResponseWriter, r *http.Request) {
    output := "create test services \n"

    // excute mehtod of object
    output += m.Create("something awesome!")
    // output += t.Create()
    output += "\n"

    w.Write([]byte(output))
}
