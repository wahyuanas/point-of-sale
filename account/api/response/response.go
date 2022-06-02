package response

import "github.com/wahyuanas/point-of-sale/account/entity"

type CommonResponse struct {
	Status  bool
	Code    int64
	Message string
}

type SignUpResponse struct {
	CommonResponse
	ID int64
}

type SignInResponse struct {
	CommonResponse
	User entity.User
}

type SignOutResponse struct {
	CommonResponse
	ID int64
}

type UpdateResponse struct {
	CommonResponse
	User entity.User
}

type DeleteResponse struct {
	CommonResponse
	ID int64
}

type GetAccountResponse struct {
	CommonResponse
	User entity.User
}

type GetAccountsResponse struct {
	User []entity.User
}
