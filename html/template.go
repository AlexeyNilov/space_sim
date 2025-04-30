package html

import (
	"net/http"
	"path/filepath"
	"sync"
	"html/template"
)

// Represents a single template
type TemplateHandler struct {
	once     sync.Once
	Filename string
	tpl    *template.Template
	Side int
	Grid string
}

// ServeHTTP handles the HTTP request.
func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.tpl = template.Must(template.ParseFiles(filepath.Join("template", t.Filename)))
	})

	// Generate the HTML content
	data := BoardData{
		Side: t.Side,
		Grid: template.HTML(t.Grid), // Convert grid to template.HTML
	}

	_ = t.tpl.Execute(w, data)
}
