package handler

import "net/http"

type UserHandlers interface {
	GetProfileHandler(w http.ResponseWriter, req *http.Request)
	GetProfilePremiumHandler(w http.ResponseWriter, req *http.Request)
}
