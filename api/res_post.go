package api

import (
    _h "github.com/javinc/acube/api/helper"
)

type Post struct {
   	Mysql
}

func (o *Post) Find() {
	_h.ConsoleLog("Post Find")		
}

func (o *Post) Get() {
	_h.ConsoleLog("Post Get")		
}

func (o *Post) Post() {
	_h.ConsoleLog("Post Post")		
}

func (o *Post) Put() {
	_h.ConsoleLog("Post Put")		
}

func (o *Post) Delete() {
	_h.ConsoleLog("Post Delete")		
}