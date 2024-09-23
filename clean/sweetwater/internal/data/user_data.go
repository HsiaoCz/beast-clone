package data

import "context"

type UserDataInter interface {
	CreateUser(context.Context)
}
