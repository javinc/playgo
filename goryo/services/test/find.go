package test

import (
    "net/http"

    "github.com/javinc/playgo/goryo/resources/test"
)

func Find(m test.Model, w http.ResponseWriter, r *http.Request) {
    output := "find test services \n"

    // excute mehtod of object
    output += m.Find("something awesome!")
    output += "\n"

    w.Write([]byte(output))
}

// this will use maps for input
// it will be losely
func FindByValue(v interface{}) string {
    return ""
}

// this will use model to get properties
// and it will be strict because of type
func FindByModel() string {
    return ""
}
