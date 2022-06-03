package main

import "github.com/wahyuanas/point-of-sale/account/delivery/grpc/client"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	accountClient *client.Client
}

func NewGraphQLServer(accountUrl string) (*Resolver, error) {
	accountClient, err := client.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		accountClient: &accountClient,
	}, nil
}
