package main

import (
	"ecom-product/config"
	"ecom-product/endpoint"
	"ecom-product/middleware"
	"ecom-product/server"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	config.LoadConfig()

	dbConnectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
	)

	dbCon, err := server.Init(dbConnectionString)
	if err != nil {
		log.Fatalf("Error initializing DB connection: %s", err)
	}

	server.DBConn = dbCon

	defer server.DBConn.Close()

	if err = server.RunMigrations(server.DBConn, "./sql_migration"); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	log.Fatal(http.ListenAndServe(":8082", controller()))

}

func controller() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.AuthMiddleware)

	r.HandleFunc("/product", endpoint.CreateProductHandler).Methods("POST")
	r.HandleFunc("/product", endpoint.GetProductsHandler).Methods("GET")
	r.HandleFunc("/product/{id}", endpoint.GetProductByIDHandler).Methods("GET")

	return r
}
