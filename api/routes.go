package api

import (
    _h "github.com/javinc/acube/api/helper"
)

type Route struct {
    Method string
    Path string
    Handler func()
}

type Routes struct {
    Route
}

func (o *Routes) SetRoutes(s interface{}) {
    _h.ConsoleLog("Routes SetRoutes")
}