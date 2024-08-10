package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
)

var Validate = validator.New()

type key int

const UserIDKey key = 0

func GetUserIDFromContext(ctx context.Context) (int, error) {
	fmt.Println(ctx)
	value := ctx.Value(UserIDKey)
	userID, ok := value.(int)
	if !ok {
		// Try other types
		if floatID, ok := value.(float64); ok {
			return int(floatID), nil
		}
		if strID, ok := value.(string); ok {
			return strconv.Atoi(strID)
		}
		return 0, errors.New("user ID not found in context or of unexpected type")
	}

	return userID, nil
}

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
