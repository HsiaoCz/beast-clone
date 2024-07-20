package service

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/anne/types"
	"github.com/sirupsen/logrus"
)

type FetchPricer interface {
	FetchPrice(context.Context, string) (*types.Price, error)
}

type DefaultPriceFetcher struct{}

var priceMap = map[string]types.Price{
	"BTH": {
		Type: "BTH",
		Min:  12.912,
		Lat:  24.867,
	},
	"GG": {
		Type: "GG",
		Min:  23.123,
		Lat:  24.987,
	},
	"DOG": {
		Type: "DOG",
		Min:  67.1324,
		Lat:  128.122,
	},
}

func (d *DefaultPriceFetcher) FetchPrice(ctx context.Context, ticker string) (*types.Price, error) {
	price, ok := priceMap[ticker]
	if !ok {
		return nil, errors.New("no record with this ticker")
	}
	logrus.WithFields(logrus.Fields{
		"RequestID": ctx.Value(types.CtxRequestID),
		"Ticker":    ticker,
		"Price":     price,
	}).Info("Get the price success")
	return &price, nil
}
