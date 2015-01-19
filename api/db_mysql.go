package api

import (
    _h "github.com/javinc/acube/api/helper"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "log"
)

const (
	db = "acube_mashdrop"
	user = "root"
	pass = "root"
)

type Mysql struct {
    Db gorm.DB
}

func (o *Mysql) Init() {
    _h.ConsoleLog("Initiating Mysql ORM")

	var err error
    
    o.Db, err = gorm.Open("mysql", user + ":" + pass + "@/" + db + "?charset=utf8&parseTime=True")
    if err != nil {
        log.Fatalf("Got error when connect database, the error is '%v'", err)
    }

    o.Db.LogMode(true)
}