package types

type Session struct {
	ID     int    `bun:"id,pk,autoincrement" json:"-"`
	UserID string `json:"userID"`
	Token  string `json:"token"`
}
