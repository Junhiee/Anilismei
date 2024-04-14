package db

import (
	"context"
	"database/sql"

	"git.virjar.com/Junhiee/anilismei/database/models"
)

type Store struct {
	*models.Queries
	db *sql.DB
}

func newStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: models.New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*models.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return  err
	}
	q := models.New(tx)
	err = fn(q)

	if err != nil {
		return err
	}
	return tx.Commit()
}