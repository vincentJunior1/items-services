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
