package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	u, _ := url.Parse("https://tabakazu.com:3000")
	fmt.Println(u.Hostname()) // => "tabakazu.com"
	fmt.Println(u.Port())     // => "3000"

	// Query String
	u, _ = url.Parse("https://tabakazu.com?family_name=taba&given_name=kazu")
	fmt.Println(u.Query().Get("family_name")) // => "taba"
	fmt.Println(u.Query().Get("given_name"))  // => "kazu"

	// Request URI
	u, _ = url.Parse("https://tabakazu.com/blogs/1")
	fmt.Println(u.RequestURI()) // => "/blogs/1"

	// URL Struct
	u, _ = url.Parse("https://tabakazu.com/admin/posts?title=hoge&body=fuga")
	fmt.Println(u.Scheme)     // => "https"
	fmt.Println(u.Host)       // => "tabakazu.com"
	fmt.Println(u.Path)       // => "/admin/posts"
	fmt.Println(u.ForceQuery) // => false
	fmt.Println(u.RawQuery)   // => "title=hoge&body=fuga"

	// Build RawQuery
	u, _ = url.Parse("https://tabakazu.com/admin")
	u.Path = path.Join(u.Path, "users/1")
	fmt.Println(u) // => "https://tabakazu.com/admin/users/1"

	// Build Query String
	u, _ = url.Parse("https://tabakazu.com?p=1")
	q := u.Query()
	q.Set("keyword", "hoge")
	u.RawQuery = q.Encode()
	fmt.Println(u) // => "https://tabakazu.com?keyword=hoge&p=1"
}
