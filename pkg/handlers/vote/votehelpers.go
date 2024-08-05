package vote

import (
	"database/sql"
	"fmt"

	"github.com/xtasysensei/go-poll/pkg/models"
)

func HasUserVoted(db *sql.DB, userID int, pollID int) (bool, error) {
	var count int
	query := `
		SELECT COUNT(*)
		FROM votes v
		INNER JOIN options o ON v.option_id = o.option_id
		WHERE v.user_id = $1 AND o.poll_id = $2
		`
	err := db.QueryRow(query, userID, pollID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CastVote(db *sql.DB, vote *models.Vote) error {
	query := `
			INSERT INTO votes(user_id, option_id) 
			VALUES ($1, $2)
			`
	_, err := db.Exec(query, vote.UserID, vote.OptionID)
	if err != nil {
		return err
	}

	return nil
}

func IsValidOptionForPoll(db *sql.DB, optionID int, pollID int) (bool, error) {
	var count int
	query := `
        SELECT COUNT(*)
        FROM options
        WHERE option_id = $1 AND poll_id = $2
    `
	err := db.QueryRow(query, optionID, pollID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("could not validate option for poll: %w", err)
	}
	return count > 0, nil
}
