package test

import (
    "errors"
    r "github.com/javinc/playgo/goryo/resources"
)

func (x *Model) Find(o Options) ([]Model, error) {
    models := []Model{}

    r.Sql.Find(&models, o.Filters)

    return models, nil
}

func (x *Model) Create(m *Model) (Model, error) {
    r.Sql.Create(&m)

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
