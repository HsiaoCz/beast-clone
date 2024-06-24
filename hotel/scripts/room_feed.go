package scripts

import (
	"context"
	"os"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoomFeed struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewRoomFeed() (*RoomFeed, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		return nil, err
	}
	return &RoomFeed{
		client: client,
		coll:   client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("ROOMCOLL")),
	}, nil
}

func (r *RoomFeed) CreateRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	// question : use which fields to keep rood unque?
	result, err := r.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = result.InsertedID.(primitive.ObjectID)
	return room, nil
}
