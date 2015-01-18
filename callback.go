package main

import (
    _h "github.com/javinc/acube/api/helper"
)

var data string

func response() {
    _h.ConsoleLog("responsing... ")

    return 
}

func request(v func()) {
    _h.ConsoleLog("requesting... ")
    
    data = "not really"
    v()

    response()
}

func main() {
    request(func () {
        _h.ConsoleLog("callback return " + data)
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