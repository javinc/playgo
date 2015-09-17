package resources

import (
    "github.com/javinc/playgo/goryo/resources/test"
    "github.com/javinc/playgo/goryo/resources/user"

    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

db, err := gorm.Open("mysql", "root:root@/ejr?charset=utf8&parseTime=True&loc=Local")

// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
db.DB()

// Then you could invoke `*sql.DB`'s functions with it
db.DB().Ping()
db.DB().SetMaxIdleConns(10)
db.DB().SetMaxOpenConns(100)

// Disable table name's pluralization
db.SingularTable(true)
db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
db.AutoMigrate(&test.Model{}, &user.Model{})
