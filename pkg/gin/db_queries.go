package gin

import (
	"fmt"
	"github.com/EraldCaka/goplater/pkg/dir"
	"os"
	"path/filepath"
)

func CreateUserQueries(username, repoName, directory string) error {
	if err := dir.CreateDir(directory); err != nil {
		return err
	}
	filePath := filepath.Join(directory, "user_queries.go")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating user_queries.go file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf(`package db

import (
	"context"
	"fmt"
	"github.com/%s/%s/types"
	"log"
)

`, username, repoName) + "\n\nfunc (pg *Postgres) GetUserByID(ctx context.Context, userID string) *types.User {\n\tquery := fmt.Sprintf(\"SELECT * FROM public.users WHERE id = '%v'\", userID)\n\trow := pg.db.QueryRow(ctx, query)\n\tvar user types.User\n\n\terr := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)\n\tif err != nil {\n\t\tlog.Printf(\"Error scanning user data: %v\\n\", err)\n\t\treturn &types.User{}\n\t}\n\n\treturn &user\n\n}\nfunc (pg *Postgres) CreateUser(ctx context.Context, u *types.UserCreate) (int, error) {\n\tquery := fmt.Sprintf(\"INSERT INTO public.users (username, password, email) VALUES ('%v','%v','%v') RETURNING id\", u.Username, u.Password, u.Email)\n\tvar userID int\n\terr := pg.db.QueryRow(ctx, query).Scan(&userID)\n\tif err != nil {\n\t\tlog.Printf(\"Unable to insert user: %v\\n\", err)\n\t\treturn 0, err\n\t}\n\treturn userID, nil\n}\nfunc (pg *Postgres) GetUsers(ctx context.Context) ([]types.User, error) {\n\tvar users []types.User\n\n\tquery := \"SELECT * FROM public.users\"\n\trows, err := pg.db.Query(ctx, query)\n\tif err != nil {\n\t\tlog.Printf(\"Error querying users: %v\\n\", err)\n\t\treturn nil, err\n\t}\n\tdefer rows.Close()\n\n\tfor rows.Next() {\n\t\tvar repo types.User\n\t\terr := rows.Scan(&repo.ID, &repo.Username, &repo.Password, &repo.Email)\n\t\tif err != nil {\n\t\t\tlog.Printf(\"Error scanning user row: %v\\n\", err)\n\t\t\tcontinue\n\t\t}\n\t\tusers = append(users, repo)\n\t}\n\n\tif err := rows.Err(); err != nil {\n\t\tlog.Printf(\"Error iterating over user rows: %v\\n\", err)\n\t\treturn nil, err\n\t}\n\n\treturn users, nil\n}\n")
	if err != nil {
		fmt.Println("Error writing to user_queries.go file:", err)
		return err
	}
	fmt.Println("user_queries.go file created successfully")
	return nil
}
