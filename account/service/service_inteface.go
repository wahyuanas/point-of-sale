package service

type AccountService interface {
	SignIn()
	SignUp()
	SignOut()
	Update()
	Delete()
	GetAccount()
	GetAccounts()
}
