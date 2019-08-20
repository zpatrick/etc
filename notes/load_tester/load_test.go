package main

import (
	"database/sql"
	"testing"
)

func TestLoadFoo(t *testing.T) {
	m := sqltest.LoadTester{
		MaxConnections: 1,
		MaxWorkers: 2,
		Connect: func() (*sql.DB, error) {
			return sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
		},
		Preload: func(db *sql.DB) error {
			return nil
		},
		Queries: []sqltest.Query{
			{
				Name: "get user",
				// todo: some notion on how frequently this query runs during the test
				// could also be handled by Run, or a Before/After which does channel ops
				Schedule: nil,
				Run: func(db *sql.DB) error {
					// todo: how would I add/get metadata so I can do CRUD on existing data?
					rows, err := db.Query("select id, name from users where id = ?", 1)
					if err != nil {
						return error
					}

					defer rows.Close()
					for rows.Next() {
						if err := rows.Scan(&id, &name); err != nil {
							return err
						}
					}

					return rows.Err()
				},
			},
		},
	}

	m.Publish(os.Stdout)
}
