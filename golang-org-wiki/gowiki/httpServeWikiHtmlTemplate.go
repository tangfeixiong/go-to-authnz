package main

import (
    "fmt"
    "html/template"
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    //fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    /*
    fmt.Fprintf(w, "<h1>Editing %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>", p.Title, p.Title, p.Body)
    */
    renderTemplate(w, "edit", p)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    //http.HandleFunc("/save/", saveHandler)
    fmt.Println("starting http service ...")
    http.ListenAndServe(os.Getenv("IP") + ":" + os.Getenv("PORT"), nil)
}
    