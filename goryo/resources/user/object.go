package user

// object model
type Model struct {
    id string `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
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
