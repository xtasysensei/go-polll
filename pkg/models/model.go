package models

import "time"

// Define your models here: type struct
type User struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
type RegisterUserPayload struct {
	Username        string `json:"username" validate:"required,alphanum"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=3,max=130"`
	ConfirmPassword string `json:"confirmpassword" validate:"required,min=3,max=130"`
}
type LoginUserPayload struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required"`
}

type Poll struct {
	PollID      int       `json:"poll_id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Options     []Option  `json:"options"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreatePollPayload struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Options     []Option `json:"options" validate:"required,dive,required"`
}

type Option struct {
	OptionID int    `json:"option_id"`
	PollID   int    `json:"poll_id"`
	Text     string `json:"text"`
}

type Vote struct {
	VoteID    int       `json:"vote_id"`
	UserID    int       `json:"user_id"`
	OptionID  int       `json:"option_id"`
	Timestamp time.Time `json:"timestamp"`
}
