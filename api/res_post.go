/*
    Post resource
*/

package api

import (
    _h "github.com/javinc/acube/api/helper"
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    "time"
)

type Posts struct {
    Id int64 `json:"id"`
    User string `sql:"size:1024" json:"user"`
    Title string `sql:"size:1024" json:"title"`
    Description string `sql:"size:1024" json:"description"`
    Active bool `json:"active"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    DeletedAt time.Time `json:"-"`
}

type PostService struct {
    Mysql
}

func (o *PostService) Migrate() {
    _h.ConsoleLog("Post Migrate")

    o.Mysql.Db.AutoMigrate(&Posts{})
}

func (o *PostService) Find(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("Post Find")

    post := []Posts{}
    
    o.Mysql.Db.Find(&post)
    w.WriteJson(&post)
}

func (o *PostService) Get(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("Post Get")

    post := Posts{}

    if o.Mysql.Db.First(&post, r.PathParam("id")).Error != nil {
        rest.NotFound(w, r)

        return
    }

    w.WriteJson(&post)
}

func (o *PostService) Post(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("Post Post")

    post := Posts{}

    if err := r.DecodeJsonPayload(&post); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        
        return
    }
    
    if err := o.Mysql.Db.Save(&post).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        
        return
    }

    w.WriteJson(&post)
}

func (o *PostService) Put(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("Post Put")

    post := Posts{}
    if o.Mysql.Db.First(&post, r.PathParam("id")).Error != nil {
        rest.NotFound(w, r)

        return
    }

    updated := Posts{}
    if err := r.DecodeJsonPayload(&updated); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)

        return
    }

    updated.Id = post.Id
    post = updated

    if err := o.Mysql.Db.Save(&post).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)

        return
    }

    w.WriteJson(&post)    
}

func (o *PostService) Delete(w rest.ResponseWriter, r *rest.Request) {
    _h.ConsoleLog("Post Delete")

    post := Posts{}
    if o.Mysql.Db.First(&post, r.PathParam("id")).Error != nil {
        rest.NotFound(w, r)
        
        return
    }

    if err := o.Mysql.Db.Delete(&post).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        
        return
    }

    w.WriteHeader(http.StatusOK)
}