package user

import "github.com/javinc/playgo/goryo/resources/user"

func Create(m *user.Model) user.Model {
    m.Name += " appended on service"

    row, _ := UserResource.Create(m)

    return row
}
