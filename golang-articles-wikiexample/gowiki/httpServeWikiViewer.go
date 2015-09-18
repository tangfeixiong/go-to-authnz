package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    title := r.URL.Path[len("/view/"):]
    fmt.Println(title)
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.ListenAndServe(os.Getenv("IP") + ":" + os.Getenv("PORT"), nil)
}
    