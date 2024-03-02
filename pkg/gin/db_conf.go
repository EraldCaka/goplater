package gin

import (
	"fmt"
	"github.com/EraldCaka/goplater/pkg/dir"
	"os"
	"path/filepath"
)

func CreateDbConfigs(username, projectName, directory string) error {
	if err := dir.CreateDir(directory); err != nil {
		return err
	}
	filePath := filepath.Join(directory, "db.go")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating db.go file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("package db\n\nimport (\n\t\"context\"\n\t\"github.com/%s/%s/util\"\n\t\"github.com/jackc/pgx/v5/pgxpool\"\n\t\"log\"\n\t\"sync\"\n)\n\ntype Postgres struct {\n\tdb *pgxpool.Pool\n}\n\nvar (\n\tpgInstance *Postgres\n\tpgOnce     sync.Once\n)\n\nfunc NewPGInstance(ctx context.Context) (*Postgres, error) {\n\tpgOnce.Do(func() {\n\t\tdb, err := pgxpool.New(ctx, util.DB_URL)\n\t\tif err != nil {\n\t\t\tlog.Println(\"Unable to connect to Postgres Db: %w\", err)\n\t\t\treturn\n\t\t}\n\t\tpgInstance = &Postgres{db}\n\t})\n\n\treturn pgInstance, nil\n}\n\nfunc (pg *Postgres) Ping(ctx context.Context) error {\n\treturn pg.db.Ping(ctx)\n}\n\nfunc (pg *Postgres) Close() {\n\tpg.db.Close()\n}", username, projectName))
	if err != nil {
		fmt.Println("Error writing to db.go file:", err)
		return err
	}
	fmt.Println("db.go file created successfully")
	return nil
}
