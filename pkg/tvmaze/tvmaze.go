package tvmaze

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/essce/flix/pkg/flix"
)

const (
	showURL = "http://api.tvmaze.com/singlesearch/shows?"
)

type Maze struct {
	client *http.Client
}

func New() *Maze {
	client := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 10 * time.Second,
		},
	}
	return &Maze{
		&client,
	}
}
func (m *Maze) Get(ctx context.Context, term string) (*flix.Show, error) {
	p := url.Values{}
	p.Set("q", term)
	p.Set("embed", "episodes")

	req, err := http.NewRequest("GET", showURL+p.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	resp, err := m.client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return nil, err
	}

	var res flix.Show
	err = json.NewDecoder(resp.Body).Decode(&res)
	return &res, err
}
