package user

import (
    "net/http"

    "github.com/javinc/playgo/goryo/resources/user"
)

func Find(m user.Model, w http.ResponseWriter, r *http.Request) {
    output := "find user services \n"

    // excute mehtod of object
    output += m.Find("something awesome!")
    // output += t.Create()
    output += "\n"

    w.Write([]byte(output))
}
