package handler

import (
	"net/http"
	"strconv"

	"golang-for-devops/internal/models"
)

func ParseBookQuery(
	r *http.Request,
) models.BookQuery {

	query := models.BookQuery{

		Title: r.URL.Query().Get("title"),

		Author: r.URL.Query().Get("author"),

		Sort: r.URL.Query().Get("sort"),
	}

	query.Page = 1
	query.Size = 10

	if value := r.URL.Query().Get("page"); value != "" {

		if page, err := strconv.Atoi(value); err == nil {

			query.Page = page
		}
	}

	if value := r.URL.Query().Get("size"); value != "" {

		if size, err := strconv.Atoi(value); err == nil {

			query.Size = size
		}
	}

	return query
}
