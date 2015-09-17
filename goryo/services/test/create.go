package test

func Create(m *Model) Model {
    row, _ := TestResource.Create(&Model{
        Name: m.Name,
    })

    return row
}
