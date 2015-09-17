package user

import (
    "net/http"

    "github.com/javinc/playgo/goryo/resources/user"
)

func Create(m user.Model, w http.ResponseWriter, r *http.Request) {
    output := "create user services \n"

    // excute mehtod of object
    output += m.Create("something awesome!")
    output += "\n"

    w.Write([]byte(output))
}
