package multi

import (
    _h "github.com/javinc/acube/api/helper"
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
    Mysql Database
}

func (o *User) Find() {
    _h.ConsoleLog("User Find")

    // get user using the schema
    user := []Users{}
    
    o.Mysql.DB.Find(&user)

    _h.ConsoleLog(user)
}

func (o *User) Get() {
    _h.ConsoleLog("User Get")
}

func (o *User) User() {
    _h.ConsoleLog("User User")
}

func (o *User) Put() {
    _h.ConsoleLog("User Put")
}

func (o *User) Patch() {
    _h.ConsoleLog("User Patch")
}

func (o *User) Delete() {
    _h.ConsoleLog("User Delete")
}