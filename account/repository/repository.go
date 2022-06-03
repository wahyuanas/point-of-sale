package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	objectvalue "github.com/wahyuanas/point-of-sale/account/api/object-value"
	"github.com/wahyuanas/point-of-sale/account/api/response"
)

type accountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{DB: db}
}

// func (acc *accountRepository) Store() error {
// 	return nil
// }
// func (acc *accountRepository) Update()        {}
// func (acc *accountRepository) Delete()        {}
//func (acc *accountRepository) GetById()       {}
// func (acc *accountRepository) GetByUserName()       {}
func (acc *accountRepository) SignIn(cmd *objectvalue.SignIn) (*response.SignInResponse, error) {
	return nil, nil
}
