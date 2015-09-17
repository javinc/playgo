package test

import "github.com/javinc/playgo/goryo/resources/test"

func Create(m *test.Model) test.Model {
    m.Name += " appended on service"

    row, _ := TestResource.Create(m)

    return row
}
