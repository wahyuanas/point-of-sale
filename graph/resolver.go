package main

import (
	"log"

	"github.com/wahyuanas/point-of-sale/account/delivery/grpc/client"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	accountClient *client.Client
}

func NewGraphQLServer(accountServiceUrl string) (*Resolver, error) {
	a, err := client.NewClient(accountServiceUrl)
	log.Println("CLIENT GRPC", accountServiceUrl, a, err)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		a,
	}, nil
}
