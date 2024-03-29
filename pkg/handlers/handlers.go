package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/xatta-trone/go-project-1/pkg/config"
	"github.com/xatta-trone/go-project-1/pkg/models"
	"github.com/xatta-trone/go-project-1/pkg/render"
)

// repository used by the handlers
var Repo *Repository

// repository type
type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, h *http.Request) {
	
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello"

	remoteIP := m.App.Session.GetString(h.Context(),"remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func Divide(w http.ResponseWriter, h *http.Request) {
	divide, err := divider(2, 0)
	if err != nil {
		_, _ = fmt.Fprint(w, err.Error())
		return
	}
	_, _ = fmt.Fprintf(w, "total sum %f", divide)

}

func divider(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("can not divide by 0")
		return 0, err
	}

	return x / y, nil

}

func AddValues(x, y int) (int, error) {
	sum := x + y

	return sum, nil
}
