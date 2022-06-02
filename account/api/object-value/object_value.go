package objectvalue

type SignUp struct {
	UserName    string
	Name        string
	Password    string
	Email       string
	PhoneNumber string
}

type SignIn struct {
	UserName string
	Password string
}

type SignOut struct {
	ID int64
}

type Update struct {
	ID          int64
	UserName    string
	Name        string
	Password    string
	Email       string
	PhoneNumber string
}

type Delete struct {
	ID int64
}

type GetAccount struct {
	ID int64
}

type GetAccounts struct {
}
