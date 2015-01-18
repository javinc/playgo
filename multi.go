package main

import (
    _h "github.com/javinc/acube/api/helper"
    multi "./multi"
)

func main() {
	_h.ConsoleLog("running main")

	// trying getting resource using GORM
	
    // init db here
	var mysql multi.Database

	// connecion
	mysql.Init()

	// resources
	user := multi.User{mysql}
	post := multi.Post{mysql}

	user.Find()
	user.Get()
	post.Get()
}