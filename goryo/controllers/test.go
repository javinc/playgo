package controllers

import (
    "log"
    "net/http"
    "encoding/json"

    "github.com/javinc/playgo/goryo/services/test"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("controllers TestHandler")

    // extract field and data
    o := new(test.Options)
    decoder.Decode(o, r.URL.Query())
    log.Println("options", o)

    t := test.Find(o)
    b, err := json.Marshal(t)
    if err != nil {
        log.Panic(err)
    }

    // check data


    // use service

    // render some value
    w.Write(b)
}
