package controller_home

import (
	"net/http"

	view_home "github.com/fbriansyah/seroja-go/internal/view/home"
	"github.com/fbriansyah/seroja-go/port/http/controller"
)

// Index implements controller.IHomeController.
func (h *HomeController) Index(w http.ResponseWriter, r *http.Request) error {
	return controller.Render(r, w, view_home.Index())
}
