package sqlite

import (
	"fmt"
	"secret-santa/pkg"

	"github.com/jmoiron/sqlx"
)

type AuthSqlite struct {
	db *sqlx.DB
}

func NewAuthSqlite(db *sqlx.DB) *AuthSqlite {
	return &AuthSqlite{db: db}
}

func (r *AuthSqlite) IsUserExists(tgUserId int) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM $1 WHERE tg_user_id = $2", usersTable, tgUserId)
}

func (r *AuthSqlite) CreateUser(user pkg.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash) values ($1, $2, $3) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
