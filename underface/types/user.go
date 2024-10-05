package types

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Synopsis string `json:"synopsis"`
	Avatar   string `json:"avatar"`
}
