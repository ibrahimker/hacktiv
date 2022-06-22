package entity

type Status struct {
	Water          int    `json:"water"`
	Wind           int    `json:"wind"`
	StatusCompiled string `json:"status_compiled`
}
type WebData struct {
	Status Status `json:"status"`
}
