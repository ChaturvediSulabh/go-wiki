package main

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.Title+".html", p.Body, 0600)
}

func main() {
}
