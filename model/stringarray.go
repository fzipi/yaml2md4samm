package model

import (
	"database/sql/driver"
	"errors"
	"strings"
)

// StringArray is a multi line string
type StringArray []string

// Scan will get a proper field for gorm
func (s *StringArray) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("failed to scan StringArray field - source is not a string")
	}
	*s = strings.Split(str, "__,__")
	return nil
}

// Value will get a value
func (s StringArray) Value() (driver.Value, error) {
	if s == nil || len(s) == 0 {
		return "{}", nil
	}
	return strings.Join(s, "__,__"), nil
}
