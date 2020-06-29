package main

import (
	"fmt"
	"net/url"
)

func main() {
	parsedUrl, _ := url.Parse("https://tabakazu.com")
	fmt.Println(parsedUrl.Hostname())
}
