package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
}


type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// Takes a context and a callback function as input, starts a new database transaction, creat a new Queries object and with that transaction and call the callback function with the created Queries object and finally commit or rollback the transaction based on the error returned by the callback function.
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	querysObject := New(tx)
	txErr := fn(querysObject)
	if txErr != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			//this means that the rollback failed
			return fmt.Errorf("transaction Error: %v, Rollback Error: %v", txErr, rbErr)
		}
		return txErr //IF the rollback is successful, return the originall transaction error
	}
	//If all operations are successful, commit the transaction
	return tx.Commit()
}