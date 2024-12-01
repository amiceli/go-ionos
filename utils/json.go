package Utils

import (
	"encoding/json"
	"io"
)

func ParseJSON[T any](body io.ReadCloser) (T, error) {
	var result T
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&result)
	return result, err
}
