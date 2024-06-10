package scripts

import (
	"context"
	"testing"

	"github.com/HsiaoCz/beast-clone/twitter/types"
)

func TestCreateUser(t *testing.T) {
	createUserParams := []types.CreateUserParams{
		{FirstName: "Hsiao", LastName: "LoveYi", NickName: "HsiaoCz", Email: "shawdml5@gmail.com", PhoneNumber: "12344556644"},
		{FirstName: "Zhang", LastName: "sanshi", NickName: "Anditdis", Email: "zhansgani@gmail.com", PhoneNumber: "13323455644"},
		{FirstName: "Hsiao", LastName: "LoveYi", NickName: "HsiaoCz", Email: "shawdml5@gmail.com", PhoneNumber: "12344556644"},
		{FirstName: "Hsiao", LastName: "LoveYi", NickName: "HsiaoCz", Email: "shawdml5@gmail.com", PhoneNumber: "12344556644"},
		{FirstName: "Hsiao", LastName: "LoveYi", NickName: "HsiaoCz", Email: "shawdml5@gmail.com", PhoneNumber: "12344556644"},
		{FirstName: "Hsiao", LastName: "LoveYi", NickName: "HsiaoCz", Email: "shawdml5@gmail.com", PhoneNumber: "12344556644"},
		{FirstName: "Hsiao", LastName: "LoveYi", NickName: "HsiaoCz", Email: "shawdml5@gmail.com", PhoneNumber: "12344556644"},
		{FirstName: "Hsiao", LastName: "LoveYi", NickName: "HsiaoCz", Email: "shawdml5@gmail.com", PhoneNumber: "12344556644"},
		{FirstName: "Hsiao", LastName: "LoveYi", NickName: "HsiaoCz", Email: "shawdml5@gmail.com", PhoneNumber: "12344556644"},
		{FirstName: "Hsiao", LastName: "LoveYi", NickName: "HsiaoCz", Email: "shawdml5@gmail.com", PhoneNumber: "12344556644"},
	}
	testStore, err := NewTestStore()
	if err != nil {
		t.Fatal(err)
	}
	for _, params := range createUserParams {
		msg := params.Validate()
		if len(msg) != 0 {
			t.Fatal(msg)
		}
		user := types.NewUserFormParams(params)
		userResp, err := testStore.CreateUser(context.TODO(), user)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", userResp)
	}
}
