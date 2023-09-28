package db

import (
	"context"
	"database/sql"
	"fmt"
)

type TxStore struct {
	*Queries
	db *sql.DB
}

func NewTxStore(db *sql.DB) *TxStore {
	return &TxStore{
		db:      db,
		Queries: New(db),
	}
}

func (s *TxStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbError := tx.Rollback(); rbError != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbError)
		}
		return err
	}

	return tx.Commit()
}

func (s *TxStore) CreateSessionTx(ctx context.Context, arg CreateSessionParams) (GetSessionBySessionIdRow, error) {
	var result GetSessionBySessionIdRow

	err := s.execTx(ctx, func(q *Queries) error {
		session, err := q.CreateSession(ctx, arg)
		if err != nil {
			return err
		}

		user, err := q.GetUser(ctx, session.UserID)
		if err != nil {
			return err
		}

		result = GetSessionBySessionIdRow{
			SessionID: session.SessionID,
			ExpiresAt: session.ExpiresAt,
			ID:        user.ID,
			Email:     user.Email,
		}

		return nil
	})

	return result, err
}
