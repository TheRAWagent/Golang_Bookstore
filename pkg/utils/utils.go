package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}