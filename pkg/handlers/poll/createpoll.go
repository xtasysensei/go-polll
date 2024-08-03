package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/xtasysensei/go-poll/pkg/database"
	"github.com/xtasysensei/go-poll/pkg/models"
	"github.com/xtasysensei/go-poll/pkg/utils"
)

func HandleCreatePoll(w http.ResponseWriter, r *http.Request) {
	// Parse JSON payload
	var payload models.CreatePollPayload
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

	// Get user ID from context
	userID, err := utils.GetUserIDFromContext(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	// Extract option texts
	var optionTexts []string
	for _, option := range payload.Options {
		optionTexts = append(optionTexts, option.Text)
	}

	// Create Poll object
	poll := models.Poll{
		UserID:      userID,
		Title:       payload.Title,
		Description: payload.Description,
		CreatedAt:   time.Now(), // Set the creation time
	}

	// Save poll and options to database
	err = CreatePoll(database.DB, userID, &poll, optionTexts)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with the created poll and its options
	utils.WriteJSON(w, http.StatusCreated, poll)
}
