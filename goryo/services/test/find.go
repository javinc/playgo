package test

import (
    "net/http"

    "github.com/javinc/playgo/goryo/resources/test"
)

func Find(m test.Model, w http.ResponseWriter, r *http.Request) {
    output := "find test services \n"

    // excute mehtod of object
    output += m.Find("something awesome!")
    // output += t.Create()
    output += "\n"

    w.Write([]byte(output))
}
