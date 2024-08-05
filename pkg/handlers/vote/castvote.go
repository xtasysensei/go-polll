package vote

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/xtasysensei/go-poll/pkg/database"
	"github.com/xtasysensei/go-poll/pkg/models"
	"github.com/xtasysensei/go-poll/pkg/utils"
)

func HandleCastVote(w http.ResponseWriter, r *http.Request) {
	pollIDStr := chi.URLParam(r, "pollId")
	pollID, err := strconv.Atoi(pollIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid poll ID"))
		return
	}

	var payload models.CastVotePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	userID, err := utils.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}
	valid, err := IsValidOptionForPoll(database.DB, payload.OptionID, pollID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if !valid {
		http.Error(w, "Invalid option for this poll", http.StatusBadRequest)
		return
	}
	// Check if the user has already cast a vote on this option
	hasVoted, err := HasUserVoted(database.DB, userID, pollID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if hasVoted {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user has already cast a vote on this poll"))
		return
	}

	// Create Vote object
	vote := models.Vote{
		UserID:    userID,
		OptionID:  payload.OptionID,
		Timestamp: time.Now(),
	}

	// Save vote to database
	err = CastVote(database.DB, &vote)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with the created vote
	utils.WriteJSON(w, http.StatusCreated, vote)
}
