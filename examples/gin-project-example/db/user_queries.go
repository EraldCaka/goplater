package db

import (
	"context"
	"fmt"
	"github.com/EraldCaka/gin-project-example/types"
	"log"
)



func (pg *Postgres) GetUserByID(ctx context.Context, userID string) *types.User {
	query := fmt.Sprintf("SELECT * FROM public.users WHERE id = '%v'", userID)
	row := pg.db.QueryRow(ctx, query)
	var user types.User

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		log.Printf("Error scanning user data: %v\n", err)
		return &types.User{}
	}

	return &user

}
func (pg *Postgres) CreateUser(ctx context.Context, u *types.UserCreate) (int, error) {
	query := fmt.Sprintf("INSERT INTO public.users (username, password, email) VALUES ('%v','%v','%v') RETURNING id", u.Username, u.Password, u.Email)
	var userID int
	err := pg.db.QueryRow(ctx, query).Scan(&userID)
	if err != nil {
		log.Printf("Unable to insert user: %v\n", err)
		return 0, err
	}
	return userID, nil
}
func (pg *Postgres) GetUsers(ctx context.Context) ([]types.User, error) {
	var users []types.User

	query := "SELECT * FROM public.users"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying users: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var repo types.User
		err := rows.Scan(&repo.ID, &repo.Username, &repo.Password, &repo.Email)
		if err != nil {
			log.Printf("Error scanning user row: %v\n", err)
			continue
		}
		users = append(users, repo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over user rows: %v\n", err)
		return nil, err
	}

	return users, nil
}
