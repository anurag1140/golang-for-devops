package service

import (
	"context"
	"testing"

	"golang-for-devops/internal/models"
	"golang-for-devops/internal/repository"
	"golang-for-devops/internal/storage"
)

func TestAddBook(t *testing.T) {

	tests := []struct {
		name    string
		input   models.Book
		wantErr bool
	}{
		{
			name: "Valid Book",
			input: models.Book{
				ID:        1,
				Title:     "Go",
				Author:    "John",
				ISBN:      "123",
				Available: true,
			},
			wantErr: false,
		},
		{
			name:    "Missing Title",
			input:   models.Book{},
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			repo := repository.NewBookRepository()

			storage := storage.NewMockStorage()

			service := NewBookService(repo, storage)

			err := service.AddBook(
				context.Background(),
				tt.input,
			)

			if tt.wantErr && err == nil {
				t.Fatal("expected validation error")
			}

			if !tt.wantErr && err != nil {
				t.Fatal(err)
			}
		})
	}
}
