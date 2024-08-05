package poll

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/xtasysensei/go-poll/pkg/database"
	"github.com/xtasysensei/go-poll/pkg/utils"
)

func RetrievePoll(w http.ResponseWriter, r *http.Request) {
	pollIDStr := chi.URLParam(r, "pollId")
	pollID, err := strconv.Atoi(pollIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid poll ID"))
		return
	}

	poll, err := GetPollByID(database.DB, pollID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, poll)
}
