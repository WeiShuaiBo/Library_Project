package system

type User struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
	Age      int    `json:"Age"`
	Sex      string `json:"Sex"`
	Identity int    `json:"Identity"`
}
