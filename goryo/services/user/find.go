package user

import "github.com/javinc/playgo/goryo/resources/user"

func Find(o *user.Options) []user.Model {
    rows, _ := UserResource.Find(*o)

    return rows
}
