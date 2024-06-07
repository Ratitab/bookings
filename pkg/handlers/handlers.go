package handlers

import (
	"net/http"

	"github.com/Ratitab/bookings/pkg/config"
	"github.com/Ratitab/bookings/pkg/models"
	"github.com/Ratitab/bookings/pkg/render"
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
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	// fmt.Println(m.App.Session.GetString(r.Context(), "remoteIP"))
	// fmt.Println(remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, Again"

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	// fmt.Println("from AboutPages", m.App.Session.GetString(r.Context(), "remoteIP"))
	stringMap["remoteIP"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Devide(w http.ResponseWriter, r *http.Request) {
}
