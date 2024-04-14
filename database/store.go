package datebase

import (
	"context"
	"database/sql"
	"fmt"

	"git.virjar.com/Junhiee/anilismei/database/models"
)

type Store struct {
	*models.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: models.New(db),
	}
}

func (store *Store) ExecTx(ctx context.Context, fn func(*models.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := models.New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}