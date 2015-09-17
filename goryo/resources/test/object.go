package test

// object model
type Model struct {
    Id string
    Name string
    Description string
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
