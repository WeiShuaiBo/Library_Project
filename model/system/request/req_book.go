package request

type ReqBook struct {
	Id           int    `json:"Id,omitempty"`
	Name         string `json:"Name,omitempty"`
	Statue       int    `json:"Statue,omitempty"`
	Author       string `json:"Author,omitempty"`
	Publish      string `json:"Publish,omitempty"`
	Introduction string `json:"Introduction,omitempty"`
}
