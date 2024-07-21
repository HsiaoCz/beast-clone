package service

import (
	"context"

	"github.com/HsiaoCz/beast-clone/anne/types"
	"github.com/gorilla/websocket"
)

type WsPriceFetcher struct {
	ws *websocket.Conn
}

func NewWsPriceFetcher(ws *websocket.Conn) *WsPriceFetcher {
	return &WsPriceFetcher{
		ws: ws,
	}
}

func (wp *WsPriceFetcher) FetchPrice(ctx context.Context, ticker string) (*types.Price, error) {
	return nil, nil
}
