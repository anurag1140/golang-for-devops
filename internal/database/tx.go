package database

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

func SafeRollback(
	ctx context.Context,
	tx pgx.Tx,
) {
	if err := tx.Rollback(ctx); err != nil &&
		err != pgx.ErrTxClosed {

		slog.Error(
			"rollback failed",
			"error",
			err,
		)
	}
}
