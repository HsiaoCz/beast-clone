package types

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Content   string `json:"content"`
	Avatar    string `json:"avatar"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"isAdmin"`
}

func (param CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}

	return errors
}
