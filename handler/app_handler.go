package handler

import (
	"net/http"
)

type IAuthorizeHandler interface {
	Authorize(w http.ResponseWriter, r *http.Request)
}

type AppHandler struct {
	AuthorizeHandler IAuthorizeHandler
}
