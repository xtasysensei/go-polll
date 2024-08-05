package poll

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

func GetPollByID(db *sql.DB, id int) (*models.Poll, error) {
	// Define the query to get poll details, options, and vote counts
	query := `
		SELECT p.poll_id, p.user_id, p.title, p.description, p.created_at,
				o.option_id, o.poll_id, o.text,
				COALESCE(COUNT(v.vote_id), 0) AS number_of_votes
		FROM polls p
		LEFT JOIN options o ON p.poll_id = o.poll_id
		LEFT JOIN votes v ON o.option_id = v.option_id
		WHERE p.poll_id = $1
		GROUP BY p.poll_id, o.option_id
		ORDER BY o.option_id
	`

	// Prepare the statement
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var poll models.Poll
	optionsMap := make(map[int][]models.Option)

	// Iterate over the rows
	for rows.Next() {
		var option models.Option
		if err := rows.Scan(
			&poll.PollID,
			&poll.UserID,
			&poll.Title,
			&poll.Description,
			&poll.CreatedAt,
			&option.OptionID,
			&option.PollID,
			&option.Text,
			&option.NumberOfVotes,
		); err != nil {
			return nil, err
		}

		// Populate the poll options
		if option.OptionID != 0 {
			optionsMap[poll.PollID] = append(optionsMap[poll.PollID], option)
		}
	}

	// Assign the options to the poll
	if options, exists := optionsMap[poll.PollID]; exists {
		poll.Options = options
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &poll, nil
}

func GetAllPolls(db *sql.DB) (*models.Poll, error) {
	query := `
		SELECT * FROM polls
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var polls models.Poll
	for rows.Next() {
		if err := rows.Scan(
			&polls.PollID,
			&polls.UserID,
			&polls.Title,
			&polls.Description,
			&polls.CreatedAt,
		); err != nil {
			return nil, err
		}

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &polls, nil
}
