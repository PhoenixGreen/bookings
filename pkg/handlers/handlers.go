package handlers

import (
	"net/http"

	"github.com/PhoenixGreen/bookings/pkg/config"
	"github.com/PhoenixGreen/bookings/pkg/models"
	"github.com/PhoenixGreen/bookings/pkg/render"
)

// Repo the repository var used by the handlers
var Repo *Repository

// repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Newrepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Newhandlers sets the repository for the handlers
func Newhandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}