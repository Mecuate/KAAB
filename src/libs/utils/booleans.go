package utils

import (
	"fmt"
)

type any interface{}

type BoolFunc struct {
	x any
}

func Boolean(x any) *BoolFunc {
	return &BoolFunc{x}
}

func (raw_value *BoolFunc) String() bool {
	x, err := ToString(raw_value.x)

	if err != nil {
		return false
	}

	return x == "true" || x == "True" || x == "TRUE" || x == "1" || len(x) > 0
}

func ToString(value interface{}) (string, error) {
	if str, ok := value.(string); ok {
		return str, nil
	}
	return "", fmt.Errorf("value is not a string")
}
