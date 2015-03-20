package main

import (
    h "github.com/javinc/acube/api/helper"
    r "./resource"
)

func main() {
	h.ConsoleLog("Main running...")

	user := r.Add("user", new(r.User))
	asset := r.Add("asset", new(r.Asset))

	user.Post()
	asset.Post()
}