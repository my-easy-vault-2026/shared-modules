package common

type Page struct {
	Current  int `json:"current" validate:"min=1,required"`
	PageSize int `json:"pageSize" validate:"max=3000,required"`
}

type PageData[T interface{}] struct {
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	Records  T   `json:"records"`
}
