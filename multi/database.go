package multi

import (
    _h "github.com/javinc/acube/api/helper"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "log"
)

const (
	db = "test"
	user = "root"
	pass = "root"
)

type Database struct {
    DB gorm.DB
}

func (o *Database) Init() gorm.DB {
    _h.ConsoleLog("Initing REST")

	var err error
    
    o.DB, err = gorm.Open("mysql", user + ":" + pass + "@/" + db + "?charset=utf8&parseTime=True")
    if err != nil {
        log.Fatalf("Got error when connect database, the error is '%v'", err)
    }

    o.DB.LogMode(true)

    return o.DB
}

func (o *Database) InitSchema(s interface{}) {
    o.DB.AutoMigrate(&s)
}