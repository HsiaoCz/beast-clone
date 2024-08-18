package dao

import (
	"context"

	"github.com/HsiaoCz/beast-clone/msgchat/types"
)


type UserCaser interface{
   CreateUser(context.Context,*types.User)(*types.User,error)
}