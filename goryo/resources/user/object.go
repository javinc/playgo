package user

import "time"

// object model
type Model struct {
    Id int `json:"id"`
    Name string `json:"name" sql:"size:255"`
    Email string `json:"email" sql:"size:255"`
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
