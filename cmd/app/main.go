package main

import (
	"context"
	"golang-for-devops/internal/auth"
	appconfig "golang-for-devops/internal/config"
	"golang-for-devops/internal/database"
	"log"
	"log/slog"
	"net/http"
	"os"

	"golang-for-devops/internal/handler"
	"golang-for-devops/internal/middleware"
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
	log.Println(applicationCfg.DatabaseURL)
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
	db := database.Connect(applicationCfg.DatabaseURL)

	bookRepo := repository.NewPostgresBookRepository(db)

	s3Client := s3.NewFromConfig(awsCfg)

	uploader := storage.NewS3Uploader(
		s3Client,
		applicationCfg.S3Bucket,
	)

	// 5. Seed Data
	// seedBooks(repo)
	// 6. Dependency Injection
	bookService := service.NewBookService(bookRepo, uploader)

	bookHandler := handler.NewBookHandler(bookService)
	userRepo := repository.NewPostgresUserRepository(db)

	refreshRepo := repository.NewPostgresRefreshTokenRepository(db)

	authService := service.NewAuthService(
		userRepo, refreshRepo,
	)

	authHandler := handler.NewAuthHandler(authService)
	// refreshHandler := handler.NewAuthHandler(authService)
	// logoutHandler := handler.NewAuthHandler(authService)

	// 7. HTTP Server
	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("POST /login", authHandler.Login)
	mux.HandleFunc("POST /refresh", authHandler.Refresh)
	mux.HandleFunc("POST /logout", authHandler.Logout)

	mux.HandleFunc("GET /books", bookHandler.GetBooks)

	mux.HandleFunc("GET /books/{id}", bookHandler.GetBookByID)

	mux.Handle(
		"PUT /books/{id}",
		middleware.Auth(
			middleware.RequireRole(
				auth.RoleAdmin,
				auth.RoleLibrarian,
			)(
				http.HandlerFunc(bookHandler.UpdateBook),
			),
		),
	)

	mux.Handle(
		"DELETE /books/{id}",
		middleware.Auth(
			middleware.RequireRole(
				auth.RoleAdmin,
			)(
				http.HandlerFunc(bookHandler.DeleteBook),
			),
		),
	)

	// Protected route
	mux.Handle(
		"POST /books",
		middleware.Auth(
			middleware.RequireRole(
				auth.RoleAdmin,
				auth.RoleLibrarian,
			)(
				http.HandlerFunc(bookHandler.AddBook),
			),
		),
	)

	logged := middleware.Logging(mux)

	log.Println("POST /login")
	log.Println("POST /refresh")
	log.Println("POST /logout")
	log.Println("GET /books")
	log.Println("POST /books")

	slog.Info("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", logged))

}
