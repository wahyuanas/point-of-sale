package server

import (
	"context"
	"database/sql"

	objectvalue "github.com/wahyuanas/point-of-sale/account/api/object-value"
	"github.com/wahyuanas/point-of-sale/account/delivery/grpc/proto/pb"
	"github.com/wahyuanas/point-of-sale/account/repository"
	"github.com/wahyuanas/point-of-sale/account/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service service.AccountService
}

func NewGRPCServer(db *sql.DB) *grpc.Server {
	r := repository.NewAccountRepository(db)
	s := service.NewAccountService(r)
	g := grpc.NewServer()
	pb.RegisterAccountServiceServer(g, &grpcServer{
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{},
		service:                           s,
	})
	return g
}

func (s *grpcServer) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	// obj := objectvalue.SignUp{UserName: in.UserName, Name: in.Name, Password: in.Password, Email: in.Email, PhoneNumber: in.PhoneNumber}
	// _, err := s.service.SignUp(&obj)
	resp := &pb.SignUpResponse{Response: &pb.CommonResponse{
		Status: true, Code: 200, Message: "SUKSES",
	}}

	return resp, nil

}
func (s *grpcServer) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	cmd := &objectvalue.SignIn{UserName: in.UserName, Password: in.Password}
	r, err := s.service.SignIn(cmd)
	if err != nil {
		return &pb.SignInResponse{Response: &pb.CommonResponse{Status: false, Code: 404, Message: "Failed"}, User: nil}, err
	}
	u := &pb.User{Id: r.User.ID, UserName: r.User.UserName, Name: r.User.Name, Password: r.User.Password, Email: r.User.Email, PhoneNumber: r.User.PhoneNumber}
	return &pb.SignInResponse{Response: &pb.CommonResponse{Status: true, Code: 200, Message: "Success"}, User: u}, err

}
func (s *grpcServer) SignOut(context.Context, *pb.SignOutRequest) (*pb.SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (s *grpcServer) Update(context.Context, *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (s *grpcServer) Delete(context.Context, *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (s *grpcServer) GetAccount(context.Context, *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (s *grpcServer) GetAccounts(context.Context, *pb.EmptyRequest) (*pb.GetAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}
