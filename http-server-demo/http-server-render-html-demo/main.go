package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// string を template に渡す例
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		t := template.Must(template.ParseFiles("templates/hello.html.tpl"))
		s := "Taro!!!"
		if err := t.ExecuteTemplate(w, "hello.html.tpl", s); err != nil {
			log.Fatal(err)
		}
	}

	// struct を template に渡す例
	hello2Handler := func(w http.ResponseWriter, req *http.Request) {
		type helloData struct {
			FirstName string
			LastName  string
		}
		t := template.Must(template.ParseFiles("templates/hello2.html.tpl"))
		param := helloData{"Ichiro", "Suzuki"}
		if err := t.ExecuteTemplate(w, "hello2.html.tpl", param); err != nil {
			log.Fatal(err)
		}
	}

	// template に変数を使う例
	hello3Handler := func(w http.ResponseWriter, req *http.Request) {
		const tpl = `
<!DOCTYPE html>
<html>
		<h1>{{.}}<h1>
</html>
`
		t, err := template.New("webpage").Parse(tpl)
		if err != nil {
			log.Fatal(err)
		}
		s := "Hello World!!!"
		if err := t.Execute(w, s); err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hello2", hello2Handler)
	http.HandleFunc("/hello3", hello3Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
