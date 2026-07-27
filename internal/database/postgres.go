package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect(connectionString string) *pgx.Conn {

	conn, err := pgx.Connect(
		context.Background(),
		connectionString,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to PostgreSQL")

	return conn
}
