package main

import (
	"io/ioutil"
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
}
