package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mirily/shorty/internal/config"
	apphttp "github.com/mirily/shorty/internal/http"
	"github.com/mirily/shorty/internal/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found, using system environment variables")
	}

	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := storage.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	router := apphttp.NewRouter()

	fmt.Println("Shorty api started on :8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
