package main

import (
	_h "github.com/javinc/acube/api/helper"
)

type Name struct {
	first string
	last string
}

type User struct {
	name Name
	Ign Name.Ign
	username string
	email string
	age int8
	active bool
}

func main() {
	// like an object when instantiated
	var james User // instantiate with defaut values
	vincent := User{Name{"James", "Vincent"}, "vince", "vince@koko.com", 22, true} // requires all field
	koko := User{username: "koko", age: 21} // optional field using key-value pair

	// displaying the fields of the object
	_h.ConsoleLog(james)
	_h.ConsoleLog(vincent)
	_h.ConsoleLog(koko)

	// accessing fields
	vincent.showInfo()
}

func (u *User) showInfo() {
	_h.ConsoleLog("\n\n------------- Showing info of", u.name.first, u.name.last, "--------------");
	_h.ConsoleLog(u.username, "email is", u.email)
	_h.ConsoleLog("and aged is", u.age)
	_h.ConsoleLog("active?", u.active)
}

func (n *Name) Ign() {
	_h.ConsoleLog("Koko")
}