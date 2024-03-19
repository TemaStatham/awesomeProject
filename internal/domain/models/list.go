package models

type List struct {
	Meta  Meta    `json:"meta"`
	Goods []Goods `json:"goods"`
}

type Meta struct {
	Total   int `json:"total"`
	Removed int `json:"removed"`
	Limit   int `json:"limit"`
	Offset  int `json:"offset"`
}
