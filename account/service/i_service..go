package service

import (
	objectvalue "github.com/wahyuanas/point-of-sale/account/api/object-value"
	"github.com/wahyuanas/point-of-sale/account/entity"
)

type AccountService interface {
	SignIn(cmd *objectvalue.SignIn) (*entity.User, error)
	SignUp(cmd *objectvalue.SignUp) (*entity.User, error)
	SignOut(cmd *objectvalue.SignOut) (*entity.User, error)
	Update(cmd *objectvalue.Update) (*entity.User, error)
	Delete(cmd *objectvalue.Delete) (*entity.User, error)
	GetAccount(cmd *objectvalue.GetAccount) (*entity.User, error)
	GetAccounts() ([]entity.User, error)
}
