package test

func Create(m *Model) string {
    output := "create test services \n"

    // excute mehtod of object
    output += TestResource.Create("something awesome!")
    output += "\n"

    return output
}
