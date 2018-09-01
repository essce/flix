package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/essce/flix/pkg/flix"
)

func (h *Handler) Query(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	ctx := r.Context()

	term := v.Get("term")
	if term == "" {
		h.writeJSONError(w, "term must be provided", 400)
		return
	}

	show, err := h.Cache.Get(ctx, term)
	if err != nil {
		h.writeJSONError(w, "failed pull from redis", 500)
		log.Printf("error: %s", err.Error())
		return
	}

	if show != nil {
		var res flix.Show
		json.Unmarshal(show, &res)
		h.writeJSONData(w, res, 200)
		return
	}

	res, err := h.API.Get(ctx, term)
	if err != nil {
		h.writeJSONError(w, "failed make api call", 500)
		log.Printf("error: %s", err.Error())
		return
	}

	if res == nil {
		h.writeJSONError(w, "show not found", 400)
		return
	}

	data, _ := json.Marshal(res)
	err = h.Cache.Set(ctx, term, data)
	if err != nil {
		h.writeJSONError(w, "failed to save response", 500)
		log.Printf("error: %s", err.Error())
		return
	}

	h.writeJSONData(w, res, 200)
}
