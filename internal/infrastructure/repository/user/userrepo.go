package user

import (
	"context"
	"database/sql"
	"jollej/db-scout/internal/domain/user"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) List(ctx *context.Context) []user.User {

	query := `SELECT * FROM USERS`
	var res []user.User
	rows, err := u.db.QueryContext(*ctx, query)
	if err != nil {
		log.Fatalf("Error fetching users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user user.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age, &user.Created_at, &user.Updated_at, &user.Is_active); err != nil {
			log.Fatalf("Error mapping fetched users: %v", err)
		}

		res = append(res, user)
	}

	return res
}
