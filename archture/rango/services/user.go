package services

import (
    "net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("user services \n"))
}
