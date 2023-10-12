package dto

import (
	"math"

	"github.com/rohanshrestha09/patra-go/enums"
)

type Query struct {
	Page  int    `form:"page,default=1"`
	Size  int    `form:"size,default=10"`
	Sort  string `form:"sort,default=id"`
	Order string `form:"order,default=desc"`
}

type GetByIDArgs struct {
	Include map[string][]string
	Exclude []string
}

type GetArgs[T any] struct {
	Include map[string][]string
	Exclude []string
	Filter  T
}

type GetAllArgs[T any] struct {
	Include   map[enums.Association][]string
	Exclude   []string
	Filter    T
	MapFilter map[string]any
	Search    map[enums.SearchColumn]string
}

type GetAllResponse[T any] struct {
	Message     string  `json:"message"`
	Data        []T     `json:"data"`
	Count       int64   `json:"count"`
	CurrentPage int     `json:"currentPage"`
	TotalPage   float64 `json:"totalPage"`
}

type Pagination struct {
	Page  int
	Size  int
	Count int64
}

func (p *Pagination) GetPages() (int, float64) {
	currentPage := p.Page

	totalPage := math.Ceil(float64(p.Count) / float64(p.Size))

	return currentPage, totalPage
}

type GetResponseReturn[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
}

func GetResponse[T any](data T, message string) GetResponseReturn[T] {
	return GetResponseReturn[T]{data, message}
}

type ResponseReturn struct {
	Message string `json:"message"`
}

func Response(message string) ResponseReturn {
	return ResponseReturn{message}
}
