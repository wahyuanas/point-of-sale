package service

import (
	objectvalue "github.com/wahyuanas/point-of-sale/account/api/object-value"
	"github.com/wahyuanas/point-of-sale/account/api/response"
	"github.com/wahyuanas/point-of-sale/account/repository"
)

type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(acc repository.AccountRepository) AccountService {
	return &accountService{acc}
}

func (s *accountService) SignIn(cmd *objectvalue.SignIn) (*response.SignInResponse, error) {
	return s.accountRepository.SignIn(cmd)

}

// func (s *accountService) SignUp(cmd *objectvalue.SignUp) (*response.SignUpResponse, error)         {}
// func (s *accountService) SignOut(cmd *objectvalue.SignOut) (*response.SignOutResponse, error)       {}
// func (s *accountService) Update(cmd *objectvalue.Update) (*response.UpdateResponse, error)         {}
// func (s *accountService) Delete(cmd *objectvalue.Delete) (*response.DeleteResponse, error)         {}
// func (s *accountService) GetAccount(cmd *objectvalue.GetAccount) (*response.GetAccountResponse, error) {}
// func (s *accountService) GetAccounts() (*response.GetAccountsResponse, error)                          {}
