package api

import "net/http"

func (h Handler) forbidden(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
}
