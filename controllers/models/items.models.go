package models

type ReqGetItems struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Quantiy int    `json:"quantity"`
}

type ParamsGetItems struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type ResponseItems struct {
	MinPrice int     `json:"min_price"`
	MaxPrice int     `json:"max_price"`
	Items    []Items `json:"items"`
}

type Items struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
