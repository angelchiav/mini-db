package database

import "os"

type Config struct {
	Path string
}

type Database struct {
	file   *os.File
	cfg    Config
	closed bool
	// schema map[string]TableSchema
}

type TableSchema struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name string
	// Type ColumnType
}
