package database

import "os"

type Database struct {
	file     *os.File
	pageSize int
	closed   bool
}
