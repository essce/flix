package flix

type Show struct {
	Name     string `json:"name"`
	Embedded struct {
		Episodes []Episode `json:"episodes"`
	} `json:"_embedded"`
}

type Episode struct {
	Name   string `json:"name"`
	Season int    `json:"season"`
	Number int    `json:"number"`
}
