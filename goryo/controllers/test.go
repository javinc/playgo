package controllers

import (
    "log"
    "net/http"

    // "github.com/javinc/playgo/goryo/services/test"
    "github.com/gorilla/schema"
)

type Person struct {
    Id int
    Name string
    Email string
}

type Sorts struct {
    Asc []string
    Desc []string
}

type Options struct {
    Filters Person
    Fields []string
    Limits []int
    Sorts
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
    render := "controllers test \n"

    // extract field and data
    o := new(Options)
    decoder := schema.NewDecoder()
    decoder.Decode(o, r.URL.Query())

    log.Println(o)

    render += Sample(map[string]string{
        "name": o.Filters.Name,
        "email": o.Filters.Email,
    })

    // check data

    // use service

    // render some value
    w.Write([]byte(render))
}

func Sample(v map[string]string) string {
    return "my name is " + v["name"] +
        " from " + v["email"]
}
