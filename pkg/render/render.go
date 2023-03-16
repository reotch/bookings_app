package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/reotch/bookings_app/pkg/config"
	"github.com/reotch/bookings_app/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data that on every page to the RenderTemplate
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// Create a template cache
	// tc = "template cache"
	var tc map[string]*template.Template

	if app.UseCache {
		// get template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// Prof Sawler creates a buffer because he want to
	// for finer-grain error-checking
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	// Render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache creates a template cache and returns a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	/*
		These two lines of code are functionally the same. One uses `make`,
		the other says 'here is the map and it is empty; see?
		There is nothing in the curly braces.' The first version is commented
		out because 'you may as well get used to this'.
	*/
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl from ./templates;
	// return a slice of strings with all of them
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through the slice and find all the files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page) // returns last element of path

		// parse the name of the file returned to var name
		// ts = "template set"
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// get the names of all files ending with *.layout.tmpl
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// In the map, for the key `name` it is equal to the current `ts` (or, template set)
		myCache[name] = ts
	}

	return myCache, nil
}
