package webapp

type User struct {
	Id       int    `json:"_"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
