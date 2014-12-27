package main

import (
    _h "github.com/javinc/acube/api/helper"
)

type callback func(string) string

func response() {
    _h.ConsoleLog("responsing... ")

    return 
}

func request(v func()) {
    _h.ConsoleLog("requesting... ")
    
    v()

    response()
}

func main() {
    request(func () {
        _h.ConsoleLog("callback return ")
    })
}

/*
$.post{'/test', {}, 
	function post(url, data, callback) {
		var confirm = confirm();
		if(confirm) {
			var sync = callback(confirm);
		}
	}
}
*/