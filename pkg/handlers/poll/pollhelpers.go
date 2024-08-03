package main

import (
	"database/sql"

	"github.com/xtasysensei/go-poll/pkg/models"
)

func CreatePoll(db *sql.DB, userID int, poll *models.Poll, options []string) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; rollback
		}
		// else: err is nil; do nothing
	}()

	// Insert the poll
	pollQuery := "INSERT INTO polls(user_id, title, description) VALUES ($1, $2, $3) RETURNING poll_id"
	err = tx.QueryRow(pollQuery, userID, poll.Title, poll.Description).Scan(&poll.PollID)
	if err != nil {
		return err
	}

	// Insert the options and collect them
	poll.Options = make([]models.Option, 0, len(options))
	optionQuery := "INSERT INTO options(poll_id, text) VALUES ($1, $2) RETURNING option_id"
	for _, optionText := range options {
		var optionID int
		err = tx.QueryRow(optionQuery, poll.PollID, optionText).Scan(&optionID)
		if err != nil {
			return err
		}
		poll.Options = append(poll.Options, models.Option{
			OptionID: optionID,
			PollID:   poll.PollID,
			Text:     optionText,
		})
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
