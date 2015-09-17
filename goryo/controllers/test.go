package controllers

import (
    "log"
    "net/http"
    "encoding/json"

    "github.com/javinc/playgo/goryo/services/test"
    tr "github.com/javinc/playgo/goryo/resources/test"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("controllers TestHandler")

    // extract field from get params
    o := new(tr.Options)
    err := decoder.Decode(o, r.URL.Query())
    if err != nil {
        log.Panic(err)
    }

    log.Println("options", o)

    // extract field from post data
    err = r.ParseForm()
    if err != nil {
        log.Panic(err)
    }

    p := new(tr.Model)
    err = decoder.Decode(p, r.PostForm)
    if err != nil {
        log.Panic(err)
    }

    log.Println("payload", p)

    // lets base on the request type
    // use service
    var result tr.Model
    switch r.Method {
    case "GET":
        result = test.Find(o)
    case "POST":
        result = test.Create(p)
    }

    // render some value
    b, err := json.Marshal(result)
    if err != nil {
        log.Panic(err)
    }

    w.Write(b)
}
