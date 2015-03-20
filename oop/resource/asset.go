/*
    Asset resource
*/

package resource

import (
    "time"
)

// object model
type Asset struct {
    Id int64 `json:"id"`

    Name string `sql:"size:1024" json:"name"`
    Description string `sql:"size:1024" json:"description"`
    Type string `sql:"size:1024" json:"type"`
    File string `sql:"size:1024" json:"file"`
    
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt time.Time `json:"-"`
}