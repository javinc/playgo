/*
    User resource
*/

package resource

import (
    "time"
)

// object model
type User struct {
    Id int64 `json:"id"`

    Name string `sql:"size:1024" json:"name"`
    Email string `sql:"size:1024" json:"email"`
    
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt time.Time `json:"-"`
}