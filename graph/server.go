package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/wahyuanas/point-of-sale/graph/generated"
)

type AppConfig struct {
	AccountServiceUrl string `envconfig:"ACCOUNT_SERVICE_URL"`
	AppPort           string `envconfig:"APP_PORT"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg AppConfig
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("PORT", cfg.AppPort)

	r, err := NewGraphQLServer(cfg.AccountServiceUrl)
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, nil))
}
