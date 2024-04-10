package entities

type Question struct {
	Name   string   `json:"name"`
	Choice []string `json:"choice"`
}
