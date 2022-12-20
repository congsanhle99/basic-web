package handlers

import (
	"net/http"

	"github.com/tsawler/go-course/pkg/config"
	"github.com/tsawler/go-course/pkg/models"
	"github.com/tsawler/go-course/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again!"
	// send data to the template
	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
