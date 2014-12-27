package main

import (
    "encoding/json"
    _h "github.com/javinc/acube/api/helper"
)

func main() {
    // http://blog.golang.org/json-and-go
    b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
    
    var f interface{}
    json.Unmarshal(b, &f)

    m := f.(map[string]interface{})

    for k, v := range m {
        // access the value
        switch vv := v.(type) {
        case string:
            _h.ConsoleLog(k, "is string", vv)
        case int:
            _h.ConsoleLog(k, "is int", vv)
        case []interface{}:
            _h.ConsoleLog(k, "is an array:")
            for i, u := range vv {
                _h.ConsoleLog(i, u)
            }
        default:
            _h.ConsoleLog(k, "is of a type I don't know how to handle")
        }
    }

    _h.ConsoleLog("--------------------------")

    var i = make(map[string]interface{})
    i["name"] = "Javinc"
    i["age"] = 22
    i["skills"] = []string{"dash", "blink"}

   	_h.ConsoleLog(i)
}