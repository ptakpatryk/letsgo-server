package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/patrykptak/letsgo-server/internal/models"
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

	snip, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}

		return
	}

	fmt.Fprintf(w, "%+v", snip)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Requesting URL Path: %s\n", r.URL.Path)
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
