package server

import (
	"context"
	"fmt"
	"net"

	"github.com/wahyuanas/point-of-sale/account/grpc/proto/pb"
	"github.com/wahyuanas/point-of-sale/account/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service service.AccountService
}

func ServeGRPC(s service.AccountService, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	pb.RegisterAccountServiceServer(srv, &grpcServer{
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{},
		service:                           s,
	})
	reflection.Register(srv)
	return srv.Serve(lis)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {

	return &pb.PostAccountResponse{Account: &pb.Account{
		Id:   "123",
		Name: "ok",
	}}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {

	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:   "123",
			Name: "v",
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {

	accounts := []*pb.Account{}

	return &pb.GetAccountsResponse{Accounts: accounts}, nil
}
