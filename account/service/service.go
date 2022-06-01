package service

import "github.com/wahyuanas/point-of-sale/account/repository"

type accountService struct {
	accountRepository repository.AccountRepository
}

func ImplAccountService(acc repository.AccountRepository) AccountService {
	return &accountService{acc}
}

func (srv *accountService) SignIn()      {}
func (srv *accountService) SignUp()      {}
func (srv *accountService) SignOut()     {}
func (srv *accountService) Update()      {}
func (srv *accountService) Delete()      {}
func (srv *accountService) GetAccount()  {}
func (srv *accountService) GetAccounts() {}
