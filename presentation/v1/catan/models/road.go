package models

type Road struct {
	ID   string `json:"id"`
	Path *Path  `json:"path"`
}
