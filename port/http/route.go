package port_http

import (
	"embed"
	"net/http"
	"time"

	"github.com/fbriansyah/seroja-go/port/http/controller"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type RouterModule struct {
	HomeCtrl controller.IHomeController
}

func NewRoute(FS *embed.FS, rm *RouterModule) http.Handler {
	requestTimeout := time.Duration(60) * time.Second
	router := chi.NewRouter()

	// Init default middleware from chi
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Recoverer)
	router.Use(chiMiddleware.Timeout(requestTimeout))

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	// Dashboard
	router.Get("/", controller.MakeHandler(rm.HomeCtrl.Index))
	return router
}
