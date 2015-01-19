package main

import (
    _h "github.com/javinc/acube/api/helper"
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    "log"
    api "./api"
)

func main() {
	_h.ConsoleLog("Running API main")

    // init db here
	var mysql api.Mysql

	// connecion on mysql db
	mysql.Init()

	// Resources
	user := api.UserService{mysql}
	post := api.PostService{mysql}

	// handler
	handler := rest.ResourceHandler{
        EnableRelaxedContentType: true,
    }
 
 	// db migration
 	user.Migrate()
 	post.Migrate()

	// routes
    err := handler.SetRoutes(
        &rest.Route{"GET", "/users", user.Find},
        &rest.Route{"GET", "/users/:id", user.Get},
        &rest.Route{"POST", "/users", user.Post},
        &rest.Route{"PUT", "/users/:id", user.Put},
        &rest.Route{"DELETE", "/users/:id", user.Delete},

		&rest.Route{"GET", "/posts", post.Find},
        &rest.Route{"GET", "/posts/:id", post.Get},
        &rest.Route{"POST", "/posts", post.Post},
        &rest.Route{"PUT", "/posts/:id", post.Put},
        &rest.Route{"DELETE", "/posts/:id", post.Delete},        
    )

    if err != nil {
        log.Fatal(err)
    }

    log.Fatal(http.ListenAndServe(":3001", &handler))
}