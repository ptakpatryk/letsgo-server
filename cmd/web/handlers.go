package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Requesting URL Path: %s\n", r.URL.Path)
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Templates
	tmplFiles := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/home.tmpl.html",
	}

	tmpl, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Requesting URL Path: %s\n", r.URL.Path)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Dispaly a specific snippet with ID %d\n", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Requesting URL Path: %s\n", r.URL.Path)
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
