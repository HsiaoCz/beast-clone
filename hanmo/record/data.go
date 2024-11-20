package main

import "go.mongodb.org/mongo-driver/mongo"

type RecordData struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func RecordDataInit(client *mongo.Client, coll *mongo.Collection) *RecordData {
	return &RecordData{
		client: client,
		coll:   coll,
	}
}
