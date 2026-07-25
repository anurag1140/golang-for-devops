package main

import (
	"context"
	appconfig "golang-for-devops/internal/appconfig"
	"log"
	"log/slog"
	"net/http"
	"os"

	"golang-for-devops/internal/handler"
	"golang-for-devops/internal/handler/middleware"
	"golang-for-devops/internal/repository"
	"golang-for-devops/internal/service"
	"golang-for-devops/internal/storage"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	// 1. Load application configuration
	applicationCfg := appconfig.Load()

	// 2. Root context
	ctx := context.Background()

	// 3. Load AWS configuration
	awsCfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		log.Fatal(err)
	}
	slog.Info(
		"AWS configuration loaded",
		"region", awsCfg.Region,
		"bucket", applicationCfg.S3Bucket,
	)

	// 4. Infrastructure
	repo := repository.NewBookRepository()

	s3Client := s3.NewFromConfig(awsCfg)

	uploader := storage.NewS3Uploader(
		s3Client,
		applicationCfg.S3Bucket,
	)

	// 5. Seed Data
	seedBooks(repo)

	// 6. Dependency Injection
	bookService := service.NewBookService(repo, uploader)

	handler := handler.NewBookHandler(bookService)

	// 7. HTTP Server
	mux := http.NewServeMux()

	mux.HandleFunc("/books", handler.Books)

	logged := middleware.Logging(mux)

	slog.Info("Server running on :8080")

	log.Fatal(http.ListenAndServe(":8080", logged))
}
