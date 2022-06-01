package entity

type User struct {
	ID          int64  `json:"id"`
	UserName    string `json:"userName"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
