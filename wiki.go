package main

import (
	"io/ioutil"
	"log"
)

//Page to render
type Page struct {
	Title string
	Body  []byte
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

func main() {
	p1 := Page{
		"wiki",
		[]byte("<h1>Hello World!</h1>"),
	}
	p1.save()
	p2, err := load(p1.Title + ".html")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("FileName:", p2.Title)
	// fmt.Printf("Content:-\n%v", string(p2.Body))
}
