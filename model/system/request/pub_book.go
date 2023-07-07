package request

type BookPage struct {
	Page   int `json:"Page" form:"Page"`
	Limit  int `json:"Limit" form:"Limit"`
	Offset int `json:"Offset" form:"Offset"`
}
