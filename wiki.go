package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Page to render
type Page struct {
	Title string
	Body  []byte
}

func main() {
	p1 := Page{
		"wiki",
		[]byte("<h1>Hello World!</h1>"),
	}
	p1.save()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	p2, err := load("wiki.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(p2.Body))
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.Title+".html", p.Body, 0600)
}

func load(fileName string) (*Page, error) {
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	p := Page{
		Title: fileName,
		Body:  body,
	}
	return &p, nil
}
