package db

import (
	"context"

	"github.com/pepsighan/nocodepress_backend/ent"
	"github.com/pkg/errors"
)

// WithTx runs the given fn within a transaction. If [fn] returns an error, then the transaction
// is rolled back otherwise it is committed.
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}

	// Handle the panics within fn as well.
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	// Try to run the fn and rollback if it fails.
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}

	// Commit otherwise.
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}

	return nil
}
