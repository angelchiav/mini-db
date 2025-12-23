package database

import (
	"encoding/binary"
	"fmt"
	"io"
)

var magic = [4]byte{'T', 'E', 'S', 'T'}

const (
	recordVersion byte = 1

	recSchema byte = 1
	recRow    byte = 2

	colInt  byte = 1
	colText byte = 2
)

func writeU16(w io.Writer, v uint16) error {
	return binary.Write(w, binary.LittleEndian, v)
}

func writeU32(w io.Writer, v uint32) error {
	return binary.Write(w, binary.LittleEndian, v)
}

func writeU64(w io.Writer, v uint64) error {
	return binary.Write(w, binary.LittleEndian, v)
}

func writeStringU16(w io.Writer, s string) error {
	if len(s) > 0xFFFF {
		return fmt.Errorf("string too long (u16): %d", len(s))
	}
	if err := writeU16(w, uint16(len(s))); err != nil {
		return err
	}

	_, err := w.Write([]byte(s))
	return err
}

func readU16(r io.Reader) (uint16, error) {
	var v uint16
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}

func readU32(r io.Reader) (uint32, error) {
	var v uint32
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}

func readU64(r io.Reader) (uint64, error) {
	var v uint64
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}

func readStringU16(r io.Reader) (string, error) {
	n, err := readU16(r)
	if err != nil {
		return "", err
	}

	buff := make([]byte, n)

	if _, err := io.ReadFull(r, buff); err != nil {
		return "", nil
	}

	return string(buff), nil
}
