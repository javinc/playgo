package test

// object model
type Model struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
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
