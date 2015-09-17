package test

import (
    "net/http"
)

func Find(w http.ResponseWriter) {
    output := "find test services \n"

    // excute mehtod of object
    output += TestResource.Find("something awesome!")
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

func Sample(o *Options) string {
    return "my name is " + o.Filters.Name +
        " from " + o.Filters.Description
}
