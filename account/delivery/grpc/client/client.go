package client

import (
	"context"

	"github.com/wahyuanas/point-of-sale/account/delivery/grpc/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	conn          *grpc.ClientConn
	serviceClient pb.AccountServiceClient
}

func NewClient(url string) (Client, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	serviceClient := pb.NewAccountServiceClient(conn)
	return &client{conn, serviceClient}, nil
}

func (c *client) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return c.serviceClient.SignUp(ctx, in)
}
func (c *client) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	return c.serviceClient.SignIn(ctx, in)
}
func (c *client) SignOut(ctx context.Context, in *pb.SignOutRequest) (*pb.SignOutResponse, error) {
	return c.serviceClient.SignOut(ctx, in)
}
func (c *client) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return c.serviceClient.Update(ctx, in)
}
func (c *client) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return c.serviceClient.Delete(ctx, in)
}
func (c *client) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	return c.serviceClient.GetAccount(ctx, in)
}
func (c *client) GetAccounts(ctx context.Context, in *pb.EmptyRequest) (*pb.GetAccountsResponse, error) {
	return c.serviceClient.GetAccounts(ctx, in)
}

func (c *client) Close() {
	c.conn.Close()
}
