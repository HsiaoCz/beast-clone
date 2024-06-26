package data

import (
	"context"
	"time"

	"github.com/HsiaoCz/beast-clone/latency/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	client *mongo.Client
	coll   *mongo.Collection
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.Database.MongoUrl))
	if err != nil {
		return nil, cleanup, err
	}
	go func() {
		if err := client.Ping(ctx, options.Client().ReadPreference); err != nil {
			log.Fatal(err)
		}
	}()
	return &Data{
		client: client,
		coll:   client.Database(c.Database.Dbname).Collection(c.Database.Coll),
	}, cleanup, nil
}
