package entities

type Question struct {
	Name   string   `json:"name"`
	Choice []string `json:"choice"`
}

type ResultQuestion struct {
	PerfumeId   string `json:"perfumeId"`
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	Notes       Notes  `json:"notes"`
	Description string `json:"description"`
	ImgUrl      string `json:"imgUrl"`
}
