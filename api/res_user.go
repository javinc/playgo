/*
    User resource
*/

package api

import (
    _h "github.com/javinc/acube/api/helper"
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
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

type UserService struct {
    Mysql
}

func (o *UserService) Find(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("User Find")

    user := []Users{}
    
    o.Mysql.Db.Find(&user)
    w.WriteJson(&user)
}

func (o *UserService) Get(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("User Get")

    user := Users{}

    if o.Mysql.Db.First(&user, r.PathParam("id")).Error != nil {
        rest.NotFound(w, r)

        return
    }

    w.WriteJson(&user)
}

func (o *UserService) Post(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("User Post")

    user := Users{}

    if err := r.DecodeJsonPayload(&user); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        
        return
    }
    
    if err := o.Mysql.Db.Save(&user).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        
        return
    }

    w.WriteJson(&user)
}

func (o *UserService) Put(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("User Put")

    user := Users{}
    if o.Mysql.Db.First(&user, r.PathParam("id")).Error != nil {
        rest.NotFound(w, r)

        return
    }

    updated := Users{}
    if err := r.DecodeJsonPayload(&updated); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)

        return
    }

    user.Name = updated.Name
    user.Email = updated.Email

    if err := o.Mysql.Db.Save(&user).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)

        return
    }

    w.WriteJson(&user)    
}

func (o *UserService) Delete(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("User Delete")

    user := Users{}
    if o.Mysql.Db.First(&user, r.PathParam("id")).Error != nil {
        rest.NotFound(w, r)
        
        return
    }

    if err := o.Mysql.Db.Delete(&user).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        
        return
    }

    w.WriteHeader(http.StatusOK)
}