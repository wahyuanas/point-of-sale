package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/wahyuanas/point-of-sale/account/delivery/grpc/server"
)

type AppConfig struct {
	DatabasePort     string `envconfig:"DB_PORT"`
	DatabaseUser     string `envconfig:"DB_USER"`
	DatabaseHost     string `envconfig:"DB_HOST"`
	DatabasePassword string `envconfig:"DB_PASSWORD"`
	DatabaseName     string `envconfig:"DB_NAME"`
	DatabaseDriver   string `envconfig:"DB_DRIVER"`
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
	log.Println("cfg.DatabaseURL", cfg.DatabaseUser)

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)

	db, err := sql.Open(DBURL, cfg.DatabaseDriver)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", cfg.DatabaseDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", cfg.DatabaseDriver)
	}
	log.Println("Starting listening on port 8080")
	port := ":8080"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", port)

	srv := server.NewGRPCServer(db)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
