package handler

import "net/http"

type AuthHandlers interface {
	SignInHandler(w http.ResponseWriter, req *http.Request)
}
