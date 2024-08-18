package dao

import (
	"context"

	"github.com/HsiaoCz/beast-clone/msgchat/types"
)

type SessionCaser interface {
	CreateSessions(context.Context, *types.Session) (*types.Session, error)
}
