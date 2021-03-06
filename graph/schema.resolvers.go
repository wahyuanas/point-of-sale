package main

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/wahyuanas/point-of-sale/account/delivery/grpc/proto/pb"
	"github.com/wahyuanas/point-of-sale/graph/generated"
	"github.com/wahyuanas/point-of-sale/graph/model"
)

func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.SignInOutput, error) {
	pb := &pb.SignInRequest{UserName: input.Username, Password: input.Password}
	s, err := r.accountClient.SignIn(ctx, pb)

	if err != nil {
		return nil, err
	}

	o := model.SignInOutput{Respon: &model.CommonOutput{Status: s.Response.Status, Code: int(s.Response.Code), Message: &s.Response.Message}, User: &model.User{ID: int(s.User.Id), Username: s.User.UserName, Name: s.User.Name, Password: s.User.Password, Email: s.User.Email, Phonenumber: s.User.PhoneNumber}}
	return &o, nil
}

func (r *queryResolver) GetAccount(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAccounts(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
