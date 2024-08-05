package poll

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/xtasysensei/go-poll/pkg/database"
	"github.com/xtasysensei/go-poll/pkg/utils"
)

func RetrievePollByID(w http.ResponseWriter, r *http.Request) {
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

func RetrieveAllPolls(w http.ResponseWriter, r *http.Request) {
	polls, err := GetAllPolls(database.DB)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, polls)
}
