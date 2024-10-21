package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"log"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		return err
	}

	if err := json.Unmarshal(body, x); err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return err
	}

	return nil
}