package test

func (o *Model) Find(s string) string {
    return "resources test finding " + s;
}

func (o *Model) Create(s string) string {
    return "resources test creating " + s;
}

func (o *Model) Get(s string) string {
    return "resources test getting " + s;
}

func (o *Model) Update(s string) string {
    return "resources test updating " + s;
}

func (o *Model) Remove(s string) string {
    return "resources test removing " + s;
}
