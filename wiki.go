package main

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	return nil
}

func main() {

}
