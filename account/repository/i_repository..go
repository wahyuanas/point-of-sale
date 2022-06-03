package repository

import (
	objectvalue "github.com/wahyuanas/point-of-sale/account/api/object-value"
	"github.com/wahyuanas/point-of-sale/account/api/response"
)

type AccountRepository interface {
	// Store() error
	// Update()
	// Delete()
	// GetById()
	// GetByUserName()
	SignIn(cmd *objectvalue.SignIn) (*response.SignInResponse, error)
}
