package client

import (
	"context"

	"github.com/wahyuanas/point-of-sale/account/delivery/grpc/proto/pb"
)

type Client interface {
	SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error)
	SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error)
	SignOut(ctx context.Context, in *pb.SignOutRequest) (*pb.SignOutResponse, error)
	Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error)
	Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error)
	GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.GetAccountResponse, error)
	GetAccounts(ctx context.Context, in *pb.EmptyRequest) (*pb.GetAccountsResponse, error)
	Close()
}
