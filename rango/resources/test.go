package resources

// object model
type Test struct {
    Name string
    Description string
}

func (o *Test) Find(s string) string {
    return "resources test finding " + s;
}

func (o *Test) Create() string {
    return "resources test create";
}
