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
	user := api.UserService{mysql}
	// post := api.Post{mysql}

	// handler
	handler := rest.ResourceHandler{
        EnableRelaxedContentType: true,
    }
 
	// routes
    err := handler.SetRoutes(
        &rest.Route{"GET", "/users", user.Find},
        &rest.Route{"GET", "/users/:id", user.Get},
        &rest.Route{"POST", "/users", user.Post},
        &rest.Route{"PUT", "/users/:id", user.Put},
        &rest.Route{"DELETE", "/users/:id", user.Delete},
    )

    if err != nil {
        log.Fatal(err)
    }

    log.Fatal(http.ListenAndServe(":3001", &handler))
}