/*
    Resource
*/

package resource

import (
    h "github.com/javinc/acube/api/helper"
)

type Resource struct {
    Name string
    Schema interface{}
}

func (o *Resource) Get() {
    h.ConsoleLog(o.Name, "Get!")
}

func (o *Resource) Post() {
    h.ConsoleLog(o.Name, "Post!")
}

func Add(name string, schema interface{}) Resource {
    h.ConsoleLog("New Instance of", name)

    // instantiate resources
    return Resource{name, schema}
}