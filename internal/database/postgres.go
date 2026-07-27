package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func Connect(connectionString string) *pgx.Conn {

	var conn *pgx.Conn
	var err error

	for i := 1; i <= 10; i++ {

		conn, err = pgx.Connect(
			context.Background(),
			connectionString,
		)

		if err == nil {
			log.Println("Connected to PostgreSQL")
			return conn
		}

		log.Printf(
			"Waiting for PostgreSQL... (%d/10)",
			i,
		)

		time.Sleep(3 * time.Second)
	}

	log.Fatal(err)

	return nil
}
