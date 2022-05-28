package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/popopame/golang-webapp-project/pkg/config"
	"github.com/popopame/golang-webapp-project/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

var TemplateCache map[string]*template.Template

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

//NewTemplates Set template Cache in the AppConfig
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template not found")
	}

	templateData = AddDefaultData(templateData)
	err := t.Execute(w, templateData)
	if err != nil {
		log.Fatal(err)
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
