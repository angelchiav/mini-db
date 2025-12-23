package database

import (
	"fmt"
	"os"
)

func OpenOrCreateDatabase(path string) (*Database, error) {
	var file *os.File
	var err error

	if _, err = os.Stat(path); os.IsNotExist(err) {
		file, err = os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("creating database file: %w", err)
		}
	} else if err == nil {
		file, err = os.OpenFile(path, os.O_RDWR, 0644)
		if err != nil {
			return nil, fmt.Errorf("opening database file: %w", err)
		}
	} else {
		return nil, fmt.Errorf("stat database file: %w", err)
	}

	return &Database{
		file: file,
	}, nil
}

func (db *Database) Close() error {
	if db.closed {
		return nil
	}

	if err := db.file.Close(); err != nil {
		return err
	}

	db.closed = true
	return nil
}
