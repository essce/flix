package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Cache  Cache
	Client *http.Client
}

type Cache interface {
	Get(context.Context, string) ([]byte, error)
	Set(context.Context, string, []byte) error
}

func (h *Handler) HTTPHandler() http.Handler {
	r := mux.NewRouter()

	s := r.PathPrefix("/v1").Subrouter()
	s.Methods("POST").Path("/search").Name("search").HandlerFunc(h.Query)

	return r
}

func (h *Handler) writeJSONError(w http.ResponseWriter, message string, code int) {
	resp := JSONErr{JSONMsg{Message: message}}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) writeJSONData(w http.ResponseWriter, data interface{}, code int) {
	resp := JSONRes{data}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)

}

type JSONRes struct {
	Data interface{} `json:"data"`
}
type JSONErr struct {
	Error JSONMsg `json:"error"`
}

type JSONMsg struct {
	Message string `json:"message"`
}
