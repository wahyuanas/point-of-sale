package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/wahyuanas/point-of-sale/account/grpc/server"
	"github.com/wahyuanas/point-of-sale/account/repository"
	"github.com/wahyuanas/point-of-sale/account/service"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var accrRepo repository.AccountRepository
	accrRepo, err = repository.ImplAccountRepository(cfg.DatabaseURL)
	if err != nil {
		log.Println(err)
	}

	log.Println("Listening on port 8080...")
	s := service.ImplAccountService(accrRepo)
	log.Fatal(server.ServeGRPC(s, 8080))
}
