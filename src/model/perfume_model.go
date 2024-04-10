package model

import "github.com/khemmaphat/scented-secrets-api/src/entities"

type PerfumeDetail struct {
	Name        string   `json:"name"`
	Brand       string   `json:"brand"`
	Description string   `json:"description"`
	HowTo       string   `json:"howTo"`
	Notes       Notes    `json:"notes"`
	Gender      string   `json:"gender"`
	ImgUrl      string   `json:"imgUrl"`
	CosineValue float32  `json:"cosineValue"`
	PerfumeType []string `json:"perfumeType"`
}

type Notes struct {
	TopNotes    []string `json:"topNotes"`
	MiddleNotes []string `json:"middleNotes"`
	BaseNotes   []string `json:"baseNotes"`
}

type GroupNotes struct {
	Name        string          `json:"name"`
	ImgGroupUrl string          `json:"imgGroupUrl"`
	Notes       []entities.Note `json:"notes"`
}
