package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgres(conf *Config, ctx context.Context) *pgxpool.Pool {
	url := "postgres://" + conf.Postgres.Username + ":" + conf.Postgres.Password + "@" + conf.Postgres.Host + ":" + conf.Postgres.Port + "/" + conf.Postgres.Dbname
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		panic(err)
	}
	config.MaxConns = int32(conf.Postgres.MaxOpenConns)
	dbpool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		panic(err)
	}

	err = Migrate(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Postgres connection initialization finished.")
	return dbpool
}

func Migrate(connectionString string) error {
	if connectionString == "" {
		return errors.New("repository: the connString was empty")
	}

	connectionString += "?sslmode=disable"

	ex, err := os.Executable()
	if err != nil {
		return err
	}

	dir := filepath.Join(filepath.Dir(ex) + "/scripts")

	migrationsPath := filepath.Join("file://" + dir)

	m, err := migrate.New(migrationsPath, connectionString)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No migrations to apply")
			return nil
		}
	}

	return err
}
