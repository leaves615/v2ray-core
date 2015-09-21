package main

import (
	"html/template"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"path"
	"path/filepath"
)

type Page struct {
	Title string
	Body  []byte

}

var templates = make(map[string]*template.Template)

func main() {
	pacServer()
}

func pacServer() {
	//	templates, _ := template.ParseFiles("html/index.html")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/config/{fileName}", getConfig).Methods("GET")
	//	http.Handle("/pac", pacHandler)
	router.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":10086", router)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/index.html")
	error := t.Execute(w, nil)
	if error != nil {
		fmt.Errorf(error.Error())
	}
}

func pacHandler(w http.ResponseWriter, r *http.Request) {
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := path.Join("release", "config", vars["fileName"])
	path, _ := filepath.Abs(file)
	d, error := ioutil.ReadFile(path)
	if error != nil {
		fmt.Printf(error.Error())
	}
	w.Write(d)
}

func saveConfig(w http.ResponseWriter, r *http.Request) {

}
