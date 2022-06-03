package service

import (
	objectvalue "github.com/wahyuanas/point-of-sale/account/api/object-value"
	"github.com/wahyuanas/point-of-sale/account/api/response"
)

type AccountService interface {
	SignIn(cmd *objectvalue.SignIn) (*response.SignInResponse, error)
	// SignUp(cmd *objectvalue.SignUp) (*response.SignUpResponse, error)
	// SignOut(cmd *objectvalue.SignOut) (*response.SignOutResponse, error)
	// Update(cmd *objectvalue.Update) (*response.UpdateResponse, error)
	// Delete(cmd *objectvalue.Delete) (*response.DeleteResponse, error)
	// GetAccount(cmd *objectvalue.GetAccount) (*response.GetAccountResponse, error)
	// GetAccounts() (response.GetAccountsResponse, error)
}
