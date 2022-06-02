package server

import (
	"context"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/wahyuanas/point-of-sale/account/grpc/proto/pb"
	"github.com/wahyuanas/point-of-sale/account/repository"
	"github.com/wahyuanas/point-of-sale/account/service"
	"google.golang.org/grpc"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service service.AccountService
}

func NewGRPCServer() *grpc.Server {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var accRepo repository.AccountRepository
	accRepo, err = repository.ImplAccountRepository(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	svc := service.ImplAccountService(accRepo)

	gsrv := grpc.NewServer()
	pb.RegisterAccountServiceServer(gsrv, &grpcServer{
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{},
		service:                           svc,
	})
	return gsrv
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
