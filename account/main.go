package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"

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
	AppPort          string `envconfig:"APP_PORT"`
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
	log.Println("cfg.DatabaseURL", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)

	db, err := sql.Open(cfg.DatabaseDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", cfg.DatabaseDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", cfg.DatabaseDriver)
	}
	log.Println("Starting listening on port ", cfg.AppPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.AppPort))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	log.Printf("Listening on %s", cfg.AppPort)

	srv := server.NewGRPCServer(db)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
