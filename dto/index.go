package dto

type Query struct {
	Page   int    `form:"page,default=1"`
	Size   int    `form:"size,default=10"`
	Sort   string `form:"sort,default=id"`
	Order  string `form:"order,default=desc"`
	Search string `form:"search"`
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
	Include   map[string][]string
	Exclude   []string
	Search    bool
	Filter    T
	MapFilter map[string]any
}

type GetAllResponse[T any] struct {
	Message string `json:"message"`
	GetAllResult[T]
}

type GetAllResult[T any] struct {
	Data        []T     `json:"data"`
	Length      int     `json:"length"`
	Count       int64   `json:"count"`
	CurrentPage int     `json:"currentPage"`
	TotalPage   float64 `json:"totalPage"`
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
