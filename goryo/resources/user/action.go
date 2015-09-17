package user

func (o *Model) Find(s string) string {
    return "resources user finding " + s;
}

func (o *Model) Create(s string) string {
    return "resources user creating " + s;
}

func (o *Model) Get(s string) string {
    return "resources user getting " + s;
}

func (o *Model) Update(s string) string {
    return "resources user updating " + s;
}

func (o *Model) Remove(s string) string {
    return "resources user removing " + s;
}
