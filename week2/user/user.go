package user

import (
	"context"
	"database/sql"
	"errors"
)

type Store struct{
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

var ErrNotFound = errors.New("user not found")

func (s Store) QueryUserByEmail(ctx context.Context, email string) (User, error)  {
	var user User
	err := s.db.QueryRow(`select user_id from user where email = $1`, email).Scan(&user.ID)
	if err !=nil{
		// sql.ErrorNoRows means "user not found" in business logic level.
		// don't return sql.ErrorNoRows or it wrapped error because it's not a real  database error
		// should use custom error for business logic level
		if err == sql.ErrNoRows{
			return User{}, ErrNotFound
		}
		return User{}, err
	}
	return user, nil
}
