package database

import (
	"fmt"
	"os"
)

func Open(cfg Config) (*Database, error) {
	if cfg.Path == "" {
		cfg.Path = "./.mydb.db"
	}

	file, err := os.OpenFile(cfg.Path, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return nil, fmt.Errorf("open database file %q: %w", cfg.Path, err)
	}

	db := &Database{
		file:   file,
		cfg:    cfg,
		closed: false,
	}

	return db, nil
}

func (db *Database) Close() error {
	if db == nil || db.closed {
		return nil
	}
	db.closed = true

	if db.file == nil {
		return nil
	}
	if err := db.file.Close(); err != nil {
		return fmt.Errorf("close database file: %w", err)
	}

	db.file = nil
	return nil
}
