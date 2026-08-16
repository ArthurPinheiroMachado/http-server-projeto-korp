package util

import (
	"encoding/json"
	"io"
)

func JsonEncodeToWriter[T any](writer io.Writer, value T) error {
	trace := CreateErrorContext("util.JSONEncodeToWriter")

	if err := json.NewEncoder(writer).Encode(value); err != nil {
		return trace.Apply(err)
	}

	return nil
}
