package user

import (
	"net/http"

	"github.com/xtasysensei/go-poll/pkg/utils"
)

func ProtectedRoute(w http.ResponseWriter, r *http.Request) {
	// Implement your protected route logic here
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "welcome to the PROTECTED PAGE"})
}
