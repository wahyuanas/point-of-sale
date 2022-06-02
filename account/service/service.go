package service

import (
	objectvalue "github.com/wahyuanas/point-of-sale/account/api/object-value"
	"github.com/wahyuanas/point-of-sale/account/entity"
	"github.com/wahyuanas/point-of-sale/account/repository"
)

type accountService struct {
	accountRepository repository.AccountRepository
}

func ImplAccountService(acc repository.AccountRepository) AccountService {
	return &accountService{acc}
}

func (srv *accountService) SignIn(cmd *objectvalue.SignIn) (*entity.User, error) {

}
func (srv *accountService) SignUp(cmd *objectvalue.SignUp) (*entity.User, error)         {}
func (srv *accountService) SignOut(cmd *objectvalue.SignOut) (*entity.User, error)       {}
func (srv *accountService) Update(cmd *objectvalue.Update) (*entity.User, error)         {}
func (srv *accountService) Delete(cmd *objectvalue.Delete) (*entity.User, error)         {}
func (srv *accountService) GetAccount(cmd *objectvalue.GetAccount) (*entity.User, error) {}
func (srv *accountService) GetAccounts() ([]entity.User, error)                          {}
