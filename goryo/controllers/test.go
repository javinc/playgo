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

    render += test.Find(o).Name

    // check data


    // use service

    // render some value
    w.Write([]byte(render))
}
