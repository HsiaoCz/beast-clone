package main

import (
	"context"
	"log"

	v1 "github.com/HsiaoCz/beast-clone/reader/api/v1"
	"github.com/HsiaoCz/beast-clone/reader/conf"
	"github.com/HsiaoCz/beast-clone/reader/storage"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := conf.ParseConfig(); err != nil {
		log.Fatal(err)
	}
	var (
		mongoURL     = conf.Conf.App.MongoUri
		dbName       = conf.Conf.App.DBName
		userCollName = conf.Conf.App.UserColl
		port         = conf.Conf.App.Port
	)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	var (
		userColl       = client.Database(dbName).Collection(userCollName)
		userMongoStore = storage.NewMongoUserStore(client, userColl)
		store          = &storage.Store{User: userMongoStore}
		userHandler    = v1.NewUserHandler(store)
		router         = gin.Default()
		av1            = router.Group("/api/v1")
	)

	{
		av1.POST("/user", userHandler.HandleCreateUser)
		av1.POST("/user/:uid", userHandler.HandleUpdateUser)
		av1.GET("/user/:uid", userHandler.HandleGetUserByID)
		av1.DELETE("/user/:uid", userHandler.HandleDeleteUserByID)
		av1.POST("/user/login", userHandler.HandleUserLogin)

	}
	router.Run(port)
}
