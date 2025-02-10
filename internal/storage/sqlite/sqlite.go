package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"urlShortener/internal/storage"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(pathStorage string) (*Storage, error) {
	const fn = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", pathStorage)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS urls(
	id INTEGER PRIMARY KEY,
	alias TEXT NOT NULL UNIQUE,
	url TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias ON urls(alias);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}

	return &Storage{db: db}, nil
}

func (s Storage) SaveUrl(instUrl, alias string) error {
	const fn = "storage.sqlite.SaveUrl"

	stmt, err := s.db.Prepare("INSERT INTO urls(url, alias) VALUES(?, ?)")
	if err != nil {
		return fmt.Errorf("%s : %w", fn, err)
	}
	_, err = stmt.Exec(instUrl, alias)
	if err != nil {
		return fmt.Errorf("%s : %w", fn, err)
	}

	return nil
}

func (s Storage) GetUrl(alias string) (string, error) {
	const fn = "storage.sqlite.GetUrl"
	var res string
	stmt, err := s.db.Prepare("SELECT url FROM urls WHERE alias IS (?)")
	if err != nil {
		return "", fmt.Errorf("%s : %w", fn, err)
	}

	err = stmt.QueryRow(alias).Scan(&res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrUrlNotFound
		}
		return "", fmt.Errorf("%s : %w", fn, err)
	}

	return res, nil
}

func (s Storage) GetAllUrls() (map[string]string, error) {
	const fn = "storage.sqlite.GetAllUrls"
	urls := make(map[string]string)

	rows, err := s.db.Query("SELECT alias, url FROM urls")
	if err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alias, url string
		if err := rows.Scan(&alias, &url); err != nil {
			return nil, fmt.Errorf("%s : %w", fn, err)
		}
		urls[alias] = url
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}

	return urls, nil
}

func (s Storage) DeleteUrl(alias string) error {
	const fn = "storage.sqlite.DeleteUrl"
	stmt, err := s.db.Prepare("DELETE FROM urls WHERE alias IS (?)")
	if err != nil {
		return fmt.Errorf("%s : %w", fn, err)
	}

	result, err := stmt.Exec(alias)
	if err != nil {
		return fmt.Errorf("%s : %w", fn, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s : %w", fn, err)
	} else if rowsAffected == 0 {
		return fmt.Errorf("%s : %w", fn, storage.ErrAliasNotFound)
	}

	return nil
}
