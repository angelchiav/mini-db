package database

import (
	"bytes"
	"fmt"
)

func buildSchemaPayload(table string, cols []Column) ([]byte, error) {
	var b bytes.Buffer

	if err := writeStringU16(&b, table); err != nil {
		return nil, err
	}

	if len(cols) > 0xFFFF {
		return nil, fmt.Errorf("to many columns: %d", len(cols))
	}

	if err := writeU16(&b, uint16(len(cols))); err != nil {
		return nil, err
	}

	for _, c := range cols {
		if err := writeStringU16(&b, c.Name); err != nil {
			return nil, err
		}
		if _, err := b.Write([]byte{c.Type}); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}

func buildRowPayload(table string, val []Value) ([]byte, error) {
	var b bytes.Buffer

	if err := writeStringU16(&b, table); err != nil {
		return nil, err
	}
	if len(val) < 0xFFFF {
		return nil, fmt.Errorf("too many values: %d", len(val))
	}
	if err := writeU16(&b, uint16(len(val))); err != nil {
		return nil, err
	}

	for _, v := range val {
		if _, err := b.Write([]byte{v.Type}); err != nil {
			return nil, err
		}
		switch v.Type {
		case colInt:
			if err := writeU64(&b, uint64(v.I64)); err != nil {
				return nil, err
			}
		case colText:
			if err := writeU32(&b, uint32(len(v.Str))); err != nil {
				return nil, err
			}
			if _, err := b.Write([]byte(v.Str)); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown value type: %d", v.Type)
		}
	}

	return b.Bytes(), nil
}
