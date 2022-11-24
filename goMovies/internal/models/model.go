package models

type MoviesData struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   *Data  `json:"data"`
	Error  string `json:"error""`
}

type Data struct {
	Movie *Movies `json:"movie"`
}

type Movies struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Genre    string  `json:"genre"`
	Rating   float64 `json:"rating"`
	Plot     string  `json:"plot"`
	Released bool    `json:"released"`
}
