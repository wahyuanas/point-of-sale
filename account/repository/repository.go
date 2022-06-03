package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	objectvalue "github.com/wahyuanas/point-of-sale/account/api/object-value"
	"github.com/wahyuanas/point-of-sale/account/api/response"
	"github.com/wahyuanas/point-of-sale/account/entity"
)

type accountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{DB: db}
}

func (acc *accountRepository) SignIn(cmd *objectvalue.SignIn) (*response.SignInResponse, error) {
	return &response.SignInResponse{CommonResponse: response.CommonResponse{Status: true, Code: 200, Message: "Success"}, User: entity.User{ID: 123, UserName: cmd.UserName, Name: "test", Password: cmd.Password, Email: "test@test.com", PhoneNumber: "081234567890"}}, nil

	//return nil, nil
}

// func (acc *accountRepository) Store() error {
// 	return nil
// }
// func (acc *accountRepository) Update()        {}
// func (acc *accountRepository) Delete()        {}
//func (acc *accountRepository) GetById()       {}
// func (acc *accountRepository) GetByUserName()       {}
