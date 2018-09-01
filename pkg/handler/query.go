package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/essce/flix/pkg/flix"
)

const (
	showURL = "http://api.tvmaze.com/singlesearch/shows?"
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

	p := url.Values{}
	p.Set("q", term)
	p.Set("embed", "episodes")

	req, err := http.NewRequest("GET", showURL+p.Encode(), nil)
	if err != nil {
		h.writeJSONError(w, "failed to create request", 500)
		log.Printf("error: %s", err.Error())
		return
	}

	resp, err := h.Client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		h.writeJSONError(w, "failed to send request", resp.StatusCode)
		log.Printf("error: %s", err.Error())
		return
	}

	var res flix.Show
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		h.writeJSONError(w, "failed to decode response", 500)
		log.Printf("error: %s", err.Error())
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
