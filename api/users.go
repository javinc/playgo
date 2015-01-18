package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "log"
    "net/http"
    "time"
)

var (
    db = "test"
    user = "root"
    pass = "root"
)

func main() {

    api := Api{}
    api.InitDB()
    api.InitSchema()

    handler := rest.ResourceHandler{
        EnableRelaxedContentType: true,
    }
    err := handler.SetRoutes(
        &rest.Route{"GET", "/user", api.GetAllUsers},
        &rest.Route{"POST", "/user", api.PostUser},
        &rest.Route{"GET", "/user/:id", api.GetUser},
        &rest.Route{"PUT", "/user/:id", api.PutUser},
        &rest.Route{"DELETE", "/user/:id", api.DeleteUser},
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Fatal(http.ListenAndServe(":8080", &handler))
}

type User struct {
    Id int64 `json:"id"`
    Name string `sql:"size:1024" json:"name"`
    Email string `sql:"size:1024" json:"email"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    DeletedAt time.Time `json:"-"`
}

type Api struct {
    DB gorm.DB
}

func (api *Api) InitDB() {
    log.Println("\nIniting DB \n")

    var err error

    api.DB, err = gorm.Open("mysql", user + ":" + pass + "@/" + db + "?charset=utf8&parseTime=True")

    // api.DB, err = gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True")
    if err != nil {
        log.Fatalf("Got error when connect database, the error is '%v'", err)
    }
    api.DB.LogMode(true)
}

func (api *Api) InitSchema() {
    log.Println("\nIniting Schema\n")

    api.DB.AutoMigrate(&User{})
}

func (api *Api) GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
    user := []User{}
    api.DB.Find(&user)
    log.Println(user)
    w.WriteJson(&user)
}

func (api *Api) GetUser(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    user := User{}
    if api.DB.First(&user, id).Error != nil {
        rest.NotFound(w, r)
        return
    }
    w.WriteJson(&user)
}

func (api *Api) PostUser(w rest.ResponseWriter, r *rest.Request) {
    user := User{}
    if err := r.DecodeJsonPayload(&user); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := api.DB.Save(&user).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteJson(&user)
}

func (api *Api) PutUser(w rest.ResponseWriter, r *rest.Request) {

    id := r.PathParam("id")
    user := User{}
    if api.DB.First(&user, id).Error != nil {
        rest.NotFound(w, r)
        return
    }

    updated := User{}
    if err := r.DecodeJsonPayload(&updated); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    user.Name = updated.Name
    user.Email = updated.Email

    if err := api.DB.Save(&user).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteJson(&user)
}

func (api *Api) DeleteUser(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    user := User{}
    if api.DB.First(&user, id).Error != nil {
        rest.NotFound(w, r)
        return
    }
    if err := api.DB.Delete(&user).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

