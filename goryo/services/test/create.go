package test

func Create(m *Model) Model {
    m.Name += " appended on create"

    return *m
}
