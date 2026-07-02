package model

type Pagination struct {
	Token     string  `json:"token"`
	NextToken *string `json:"next_token"`
	Limit     uint16  `json:"limit"`
}
