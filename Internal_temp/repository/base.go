package Repository

import (
	db "awesomeProject/db/sqlc"
	"context"

	"database/sql"
	"errors"
)

const DbTransactionKey string = "transaction_key"

type BaseRepository struct {
	dbtx    db.DBTX
	Queries *db.Queries
	db      *sql.DB
}

func NewBaseRepository(q *db.Queries, DB *sql.DB) *BaseRepository {
	return &BaseRepository{
		dbtx:    DB,
		Queries: q,

		db: DB,
	}
}

func (b *BaseRepository) storeTransaction(ctx context.Context, tx *sql.Tx) (context.Context, error) {
	newCtx := context.WithValue(ctx, DbTransactionKey, tx)
	return newCtx, nil
}

func (b *BaseRepository) getTransaction(ctx context.Context) (*sql.Tx, error) {
	tx := ctx.Value(DbTransactionKey)
	if tx == nil {
		return nil, nil
	}
	return tx.(*sql.Tx), nil
}

func (b *BaseRepository) BeginTransaction(ctx context.Context) (context.Context, error) {
	tx, err := b.db.Begin()
	if err != nil {
		return nil, err
	}
	return b.storeTransaction(ctx, tx)
}

func (b *BaseRepository) CommitTransaction(ctx context.Context) error {
	tx, err := b.getTransaction(ctx)
	if err != nil {
		return err
	}
	if tx == nil {
		return errors.New("no transaction found on current context")
	}

	err = tx.Commit()

	if err != nil {
		return err
	}
	return nil
}

func (b *BaseRepository) RollbackTransaction(ctx context.Context) error {
	tx, err := b.getTransaction(ctx)
	//b.Queries = db.New(tx)
	if err != nil {
		return err
	}
	if tx == nil {
		return errors.New("no transaction found on current context")
	}

	err = tx.Rollback()

	if err != nil {
		return err
	}
	return nil
}

func (b *BaseRepository) GetConnection(ctx context.Context) error {
	tx, err := b.getTransaction(ctx)

	if err != nil {
		return err
	}

	if tx != nil {
		b.Queries = db.New(tx)
	} else {
		b.Queries = db.New(b.db)
	}

	return nil
}
