package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/beast-clone/heather/protopkg"
	"github.com/HsiaoCz/beast-clone/heather/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	req, err := http.NewRequest("get", c.endpoint, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("ticker", ticker)
	req.URL.RawQuery = query.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	priceResp := &types.PriceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}
	return priceResp, nil
}

func NewGRPCClient(remoteAddr string) (protopkg.PriceFetcherClient, error) {
	conn, err := grpc.NewClient(remoteAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return protopkg.NewPriceFetcherClient(conn), nil
}
