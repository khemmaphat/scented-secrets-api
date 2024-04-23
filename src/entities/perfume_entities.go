package entities

type Perfume struct {
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	HowTo       string  `json:"howTo"`
	Notes       string  `json:"notes"`
	Gender      string  `json:"gender"`
	ImgUrl      string  `json:"imgUrl"`
	CosineValue float32 `json:"cosineValue"`
}

type PerfumePaginationRequest struct {
	Search      string
	SearchType  string
	SearchGroup string
	CosineValue float32
	PageSize    int
	PageNum     int
	Gender      string
}

type PerfumeResponse struct {
	PerfumeId string `json:"perfumeId"`
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	ImgUrl    string `json:"imgUrl"`
}

type Note struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Group       string `json:"group,omitempty"`
	ImgUrl      string `json:"imgUrl,omitempty"`
}

type Notes struct {
	TopNotes    []string `json:"topNotes"`
	MiddleNotes []string `json:"middleNotes"`
	BaseNotes   []string `json:"baseNotes"`
}

type PerfumeComment struct {
	Name    string `json:"name"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

type PerfumePath struct {
	Path string `json:"path"`
}
