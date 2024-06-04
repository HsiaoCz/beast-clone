package main

import (
	"context"
	"log"
	"os"

	v1 "github.com/HsiaoCz/beast-clone/reader/api/v1"
	"github.com/HsiaoCz/beast-clone/reader/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var (
		mongoURL     = os.Getenv("MONGOURL")
		dbName       = os.Getenv("DBNAME")
		userCollName = os.Getenv("USERCOLL")
		port         = os.Getenv("PORT")
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

	}
	router.Run(port)
}
