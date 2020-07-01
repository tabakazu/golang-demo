package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://tabakazu.com/")
	fmt.Println(resp.Status)     // => "200 OK"
	fmt.Println(resp.StatusCode) // => 200
	fmt.Println(resp.Proto)      // => "HTTP/2.0"

	req := resp.Request
	fmt.Println(req.Method) // => "GET"
	fmt.Println(req.Proto)  // => "HTTP/1.1"
	fmt.Println(req.Host)   // => "tabakazu.com"
}
