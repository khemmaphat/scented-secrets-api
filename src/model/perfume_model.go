package model

import "github.com/khemmaphat/scented-secrets-api/src/entities"

type PerfumeDetail struct {
	Name        string                 `json:"name"`
	Brand       string                 `json:"brand"`
	Notes       GroupNotePerfumeDetail `json:"notes"`
	PerfumeType []string               `json:"perfumeType"`
	Description string                 `json:"description"`
	HowTo       string                 `json:"howTo"`
	Gender      string                 `json:"gender"`
	ImgUrl      string                 `json:"imgUrl"`
	CosineValue float32                `json:"cosineValue"`
}

type GroupNotePerfumeDetail struct {
	TopNotes    []entities.Note `json:"topNotes"`
	MiddleNotes []entities.Note `json:"middleNotes"`
	BaseNotes   []entities.Note `json:"baseNotes"`
}

type GroupNotes struct {
	Name        string          `json:"name"`
	ImgGroupUrl string          `json:"imgGroupUrl"`
	Notes       []entities.Note `json:"notes"`
}

type ResultMixedPerfume struct {
	Notes       entities.Notes `json:"notes"`
	Description string         `json:"description"`
}

type AveragePerfumeComment struct {
	AverageRating float32                   `json:"averageRating"`
	LikeRating    float32                   `json:"likeRating"`
	OkRating      float32                   `json:"okRating"`
	DislikeRating float32                   `json:"dislikeRating"`
	Comments      []entities.PerfumeComment `json:"comments"`
}
