package main

import (
	"io/ioutil"
	"os"
	"strings"

	migrate "github.com/rubenv/sql-migrate"
)

func parseSQLQueries(folder string) ([]string, error) {
	queries := []string{}
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	for _, fi := range files {
		fl, err := os.Open(strings.Join([]string{folder, fi.Name()}, "/"))
		if err != nil {
			return nil, err
		}
		defer fl.Close()
		b, err := ioutil.ReadAll(fl)
		if err != nil {
			return nil, err
		}
		queries = append(queries, string(b))
	}
	return queries, nil
}

func migrateT(up []string, down []string) error {
	migrations := migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{&migrate.Migration{
			Id:   "1",
			Up:   up,
			Down: down,
		},
		},
	}
	_, err := migrate.Exec(db, dbDriver, migrations, migrate.Up)
	if err != nil {
		return err
	}
	return nil
}

func runMigrations() error {
	up, err := parseSQLQueries("up")
	if err != nil {
		return err
	}
	down, err := parseSQLQueries("down")
	if err != nil {
		return err
	}
	return migrateT(up, down)
}
