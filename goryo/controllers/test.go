package controllers

import (
    "log"
    "net/http"

    "github.com/javinc/playgo/goryo/services/test"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
    render := "controllers test \n"

    // extract field and data
    o := new(test.Options)
    decoder.Decode(o, r.URL.Query())

    log.Println(o)

    // render += Sample(map[string]string{
    //     "name": o.Filters.Name,
    //     "description": o.Filters.Description,
    // })

    // check data


    // use service

    // render some value
    w.Write([]byte(render))
}

func Sample(v map[string]string) string {
    return "my name is " + v["name"] +
        " from " + v["email"]
}
