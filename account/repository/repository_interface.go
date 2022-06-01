package repository

type AccountRepository interface {
	Store() error
	Update()
	Delete()
	GetById()
	GetByUserName()
}
