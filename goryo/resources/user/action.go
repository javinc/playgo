package user

import "errors"

func (x *Model) Find(o Options) ([]Model, error) {
    Models := []Model {
        o.Filters,
        o.Filters,
    }

    return Models, nil
}

func (x *Model) Create(m *Model) (Model, error) {
    m.Name += " appended on resource"

    return *m, nil
}

func (x *Model) Get(i int) (Model, error) {
    return *x, nil
}

func (x *Model) Remove(i int) error {
    return errors.New("Remove method not available")
}

func (x *Model) Update(i int) error {
    return errors.New("Update method not available")
}
