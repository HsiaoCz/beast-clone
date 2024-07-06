package types

import "time"

type User struct {
	ID              int       `bun:"id,pk,autoincrement" json:"-"`
	UserID          string    `json:"userID"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Password        string    `json:"-"`
	Content         string    `json:"content"`
	Avatar          string    `json:"avatar"`
	Synopsis        string    `json:"synopisis"`
	Phone           string    `json:"-"`
	BackgroundImage string    `json:"backgroundImage"`
	EmailVerifiedAt time.Time `json:"-"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"-"`
}
