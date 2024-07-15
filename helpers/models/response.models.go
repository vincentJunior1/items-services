package models

type Response struct {
	Meta MetaRes     `json:"meta"`
	Data interface{} `json:"data,omitempty"`
	Page Pagination  `json:"pagination,omitempty"`
}

type MetaRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Pagination struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}
