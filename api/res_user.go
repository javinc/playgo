package api

import (
    _h "github.com/javinc/acube/api/helper"
    "github.com/ant0ine/go-json-rest/rest"
    "time"
)

type Users struct {
    Id int64 `json:"id"`
    Name string `sql:"size:1024" json:"name"`
    Email string `sql:"size:1024" json:"email"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    DeletedAt time.Time `json:"-"`
}

type User struct {
    Mysql
}

func (o *User) Find(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("User Find")

    user := []Users{}
    
    o.Mysql.DB.Find(&user)

    _h.ConsoleLog(user)

    w.WriteJson(&user)
}

func (o *User) Get(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("User Get")
}

func (o *User) Post(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("User Post")
}

func (o *User) Put(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("User Put")
}

func (o *User) Patch(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("User Patch")
}

func (o *User) Delete(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("User Delete")
}