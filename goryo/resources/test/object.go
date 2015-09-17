package test

import "time"

const tableName = "user"

// object model
type Model struct {
    Id int `json:"id"`
    Name string `json:"name" sql:"size:255"`
    Description string `json:"description" sql:"size:255"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"-"`
}

type Options struct {
    Filters Model
    Fields []string
    Limits []int
    Sorts struct {
        Asc []string
        Desc []string
    }
}

// specify table name required by ORM
func (x Model) TableName() string {
    return tableName
}
