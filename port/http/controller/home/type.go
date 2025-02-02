package controller_home

import (
	"github.com/fbriansyah/seroja-go/port/http/controller"
)

type HomeController struct{}

func New() controller.IHomeController {
	return &HomeController{}
}
