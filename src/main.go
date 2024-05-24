package main

import (
	"fmt"

	"github.com/Dwibi/beli-mang/src/drivers/db"
	"github.com/Dwibi/beli-mang/src/http"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := db.CreateConnection()

	if err != nil {
		fmt.Println("Error creating database connection:", err)
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	// Load the Shared AWS Configuration (~/.aws/config)
	// cfg, err := config.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create an Amazon S3 service client
	// client := s3.NewFromConfig(cfg)
	// uploader := manager.NewUploader(client)

	// h := http.New(&http.Http{
	// 	DB:       dbConnection,
	// 	Uploader: uploader,
	// })

	h := http.New(&http.Http{
		DB: db,
		// Uploader: uploader,
	})

	h.Launch()

}
