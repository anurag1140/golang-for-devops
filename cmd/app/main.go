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

	_ "golang-for-devops/docs"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Library API
// @version 1.0
// @description Library Management System API
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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
	loanRepo := repository.NewPostgresLoanRepository(db)

	loanService := service.NewLoanService(
		loanRepo,
	)

	loanHandler := handler.NewLoanHandler(
		loanService,
	)

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

	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

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

	mux.Handle(
		"POST /loans",
		middleware.Auth(
			middleware.RequireRole(
				auth.RoleLibrarian,
				auth.RoleAdmin,
			)(
				http.HandlerFunc(
					loanHandler.IssueBook,
				),
			),
		),
	)

	mux.Handle(
		"POST /returns/{bookId}",
		middleware.Auth(
			middleware.RequireRole(
				auth.RoleLibrarian,
				auth.RoleAdmin,
			)(
				http.HandlerFunc(
					loanHandler.ReturnBook,
				),
			),
		),
	)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	logged := middleware.Logging(mux)

	log.Println("POST /login")
	log.Println("POST /refresh")
	log.Println("POST /logout")
	log.Println("GET /books")
	log.Println("POST /books")

	slog.Info("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", logged))

}
