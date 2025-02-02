package controller

import "net/http"

type IHomeController interface {
	Index(w http.ResponseWriter, r *http.Request) error
}
