package database

import (
	"fmt"
	"io"
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

func readRecord(r io.Reader) (*Record, error) {
	var m [4]byte
	if _, err := io.ReadFull(r, m[:]); err != nil {
		return nil, err
	}
	if m != magic {
		return nil, fmt.Errorf("bad magic: %v", m)
	}

	hdr := make([]byte, 2)
	if _, err := io.ReadFull(r, hdr); err != nil {
		return nil, err
	}
	typ, ver := hdr[0], hdr[1]

	n, err := readU32(r)
	if err != nil {
		return nil, err
	}
	payload := make([]byte, n)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}

	return &Record{Type: typ, Version: ver, Payload: payload}, nil
}
