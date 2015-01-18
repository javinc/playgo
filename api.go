package main

import (
    _h "github.com/javinc/acube/api/helper"
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    "log"
    api "./api"
)

func main() {
	_h.ConsoleLog("running API main")

    // init db here
	var mysql api.Mysql

	// connecion on mysql db
	mysql.Init()

	// Resources
	user := api.User{mysql}
	// post := api.Post{mysql}

	// routes
	handler := rest.ResourceHandler{
        EnableRelaxedContentType: true,
    }
 
    err := handler.SetRoutes(
        &rest.Route{"GET", "/user", user.Find},
        &rest.Route{"GET", "/user/:id", user.Get},
        &rest.Route{"POST", "/user", user.Post},
        &rest.Route{"PUT", "/user/:id", user.Put},
        &rest.Route{"PATCH", "/user/:id", user.Patch},
        &rest.Route{"DELETE", "/user/:id", user.Delete},
    )

    if err != nil {
        log.Fatal(err)
    }

    log.Fatal(http.ListenAndServe(":3001", &handler))
}