package auth

import (
	"fmt"
	"sunset/app/db"

	"github.com/anthdm/superkit/kit"
	v "github.com/anthdm/superkit/validate"
)

var profileSchema = v.Schema{
	"firstName": v.Rules(v.Min(3), v.Max(50)),
	"lastName":  v.Rules(v.Min(3), v.Max(50)),
}

type ProfileFormValues struct {
	ID       int    `form:"id"`
	Username string `form:"username"`
	Email    string
	Success  string
}

func HandleProfileShow(kit *kit.Kit) error {
	auth := kit.Auth().(Auth)

	var user User
	err := db.Query.NewSelect().
		Model(&user).
		Where("id = ?", auth.UserID).
		Scan(kit.Request.Context())
	if err != nil {
		return err
	}

	formValues := ProfileFormValues{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return kit.Render(ProfileShow(formValues))
}

func HandleProfileUpdate(kit *kit.Kit) error {
	var values ProfileFormValues
	errors, ok := v.Request(kit.Request, &values, profileSchema)
	if !ok {
		return kit.Render(ProfileForm(values, errors))
	}

	auth := kit.Auth().(Auth)
	if auth.UserID != values.ID {
		return fmt.Errorf("unauthorized request for profile %d", values.ID)
	}
	_, err := db.Query.NewUpdate().
		Model((*User)(nil)).
		Set("username = ?", values.Username).
		Where("id = ?", auth.UserID).
		Exec(kit.Request.Context())
	if err != nil {
		return err
	}

	values.Success = "Profile successfully updated!"
	values.Email = auth.Email

	return kit.Render(ProfileForm(values, v.Errors{}))
}
