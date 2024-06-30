package types

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Content  string `json:"content"`
	Avatar   string `json:"avatar"`
}
