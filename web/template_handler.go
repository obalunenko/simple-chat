package web

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/objx"
)

// templateHandler represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	path     string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join(t.path, t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	if err := t.templ.Execute(w, data); err != nil {
		http.Error(w, "failed to execute template", http.StatusInternalServerError)
	}
}

// NewTemplateHandler creates new templateHandler object that satisfy http.Handler interface
func NewTemplateHandler(path string, filename string) http.Handler {
	return &templateHandler{
		once:     sync.Once{},
		filename: filename,
		path:     path,
		templ:    &template.Template{},
	}
}
