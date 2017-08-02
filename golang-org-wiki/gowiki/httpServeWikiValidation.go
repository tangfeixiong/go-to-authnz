package main

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "net/http"
    "regexp"
    "errors"
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

var templates = template.Must(template.ParseFiles("view.html", "edit.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    /*
    t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    */
    err := templates.ExecuteTemplate(w, tmpl + ".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var validPath = regexp.MustCompile("^/(view|edit|save)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("Invalid Page Tile")
    }
    return m[2], nil
}
    

func viewHandler(w http.ResponseWriter, r *http.Request) {
    //title := r.URL.Path[len("/view/"):]
    title, err := getTitle(w, r)
    if err != nil {
        return 
    }
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/" + title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    //title := r.URL.Path[len("/edit/"):]
    title, err := getTitle(w, r)
    if err != nil {
        return 
    }
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    //title := r.URL.Path[len("/save/"):]
    title, err := getTitle(w, r)
    if err != nil {
        return 
    }
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err = p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    fmt.Println("starting http service ...")
    http.ListenAndServe(os.Getenv("IP") + ":" + os.Getenv("PORT"), nil)
}
    