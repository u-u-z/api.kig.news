package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	value, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return string(value), nil
}

func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	result := map[string]interface{}{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}

	*j = result
	return nil
}
