package main

import (
    h "github.com/javinc/acube/api/helper"
)

type Dog struct {
	Name string
	Cute bool
}

type Dragon struct {
	Name string
	WingSpan int64
}

type Animal struct {
	Kind string
	Schema interface{}
}

func (a *Animal) Speak() {
	h.ConsoleLog("Hi,", a.Kind, "Speaking...")
}

func (a *Animal) Walk() {
	h.ConsoleLog(a.Kind, "Speaking...")
}

func main() {
	h.ConsoleLog("Main running...")

	// dog named koko
	koko := Animal{"Dog", new(Dog)}
	koko.Speak()

	// dragon named momo
	momo := Animal{"Dragon", new(Dragon)}
	momo.Speak()
}