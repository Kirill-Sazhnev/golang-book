package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type Guests []string

func main() {
	http.HandleFunc("/add", makeHandler(addHandler))
	http.HandleFunc("/edit", makeHandler(editHandler))
	http.HandleFunc("/rewrite", makeHandler(rewriteHandler))
	http.HandleFunc("/", makeHandler(indexHandler))
	log.Fatal(http.ListenAndServe(":8765", nil))
}

func (p Guests) save() error {
	guests := strings.Join(p, "\n")
	file, err := os.OpenFile("Guests.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(guests + "\n")
	return err
}

func loadGuests() (*Guests, error) {
	body, err := os.ReadFile("Guests.txt")
	list := Guests(strings.Fields(string(body)))
	if err != nil {
		return nil, err
	}
	return &list, nil
}

var validPath = regexp.MustCompile("^/([a-zA-Z0-9]*)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

// indexHandler serves the main page
func indexHandler(w http.ResponseWriter, req *http.Request) {
	list, err := loadGuests()
	if err != nil {
		http.Redirect(w, req, "/add", http.StatusFound)
		return
	}
	templateHandler(w, list, indexHTML)
}

// addHandler add a name to the names list
func addHandler(w http.ResponseWriter, req *http.Request) {
	names := req.FormValue("name")
	list := Guests(strings.Fields(names))
	err := list.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	list, err := loadGuests()
	if err != nil {
		http.Redirect(w, r, "/add", http.StatusFound)
		return
	}

	templateHandler(w, list, editHTML)
}

func rewriteHandler(w http.ResponseWriter, r *http.Request) {
	list := r.FormValue("guests")
	err := os.WriteFile("Guests.txt", []byte(list), 0600)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

var templ = template.New("HTML")

func templateHandler(w http.ResponseWriter, list *Guests, t string) {
	templ, _ = templ.Parse(t)
	err := templ.Execute(w, list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var indexHTML = `
<!DOCTYPE html>
<html>
    <head>
		<title>Guest Book ::Web GUI</title>
    </head>
    <body>
		<h1>Guest Book :: Web GUI</h1>
		<form action="/add" method="post">
		Name: <input name="name" /><submit value="Sign Guest Book">
		</form>
		<hr />
		<h4>Previous Guests</h4>
		<ul>
			{{range .}}
			<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
</html>
`
var editHTML = `
<h1>Editing Guest List</h1>

<form action="/rewrite" method="POST">
<div><textarea name="guests" rows="20" cols="80">
{{range .}}{{printf "%s" .}}
{{end}}</textarea></div>
<div><input type="submit" value="Save"></div>
</form>`
