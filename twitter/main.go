package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/twitter/app"
	"github.com/HsiaoCz/beast-clone/twitter/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			log.Fatal(err)
		}
	}()

	var (
		port             = os.Getenv("PORT")
		dbname           = os.Getenv("DBNAME")
		userCollName     = os.Getenv("USERCOLL")
		postCollName     = os.Getenv("POSTCOLL")
		commentCollName  = os.Getenv("COMMENTCOLL")
		userColl         = client.Database(dbname).Collection(userCollName)
		postColl         = client.Database(dbname).Collection(postCollName)
		commentColl      = client.Database(dbname).Collection(commentCollName)
		mongoUserCase    = db.NewMongoUserCase(client, userColl)
		mongoPostCase    = db.NewMongoPostStore(client, postColl)
		mongoCommentCase = db.NewMongoCommentStore(client, commentColl)
		dbs              = &db.DBS{Uc: mongoUserCase, Pc: mongoPostCase, CS: mongoCommentCase}
		userApp          = app.NewUserApp(dbs)
		postApp          = app.NewPostApp(dbs)
		commentApp       = app.NewCommentApp(dbs)
		router           = chi.NewMux()
	)

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	// TODO need group router
	// and use middleware auth
	router.Post("/user", app.TransferHandlerfunc(userApp.HandleCreateUser))
	router.Get("/user/{uid}", app.TransferHandlerfunc(userApp.HandleGetUserByID))
	router.Delete("/user", app.TransferHandlerfunc(userApp.HandleDeleteUserByID))
	router.Post("/user/update", app.TransferHandlerfunc(userApp.HandleUpdateUserByID))

	// TODO need group router
	// posts handlers need auth middleware
	router.Post("/posts", app.TransferHandlerfunc(postApp.HandleCreatePost))
	router.Get("/posts/{uid}", app.TransferHandlerfunc(postApp.HandleGetPostsByUserID))
	router.Delete("/posts/{pid}", app.TransferHandlerfunc(postApp.HandleDeletePost))

	// TODO need group router comment
	// posts handlers need auth middleware
	router.Post("/comments", app.TransferHandlerfunc(commentApp.HandleCreateComment))
	router.Get("/comments/{pid}", app.TransferHandlerfunc(commentApp.HandleGetCommentByPostID))
	router.Delete("/comments/{cid}", app.TransferHandlerfunc(commentApp.HandleDeleteCommentByID))

	logrus.WithFields(logrus.Fields{
		"listen address": port,
	}).Info("http server is running")
	http.ListenAndServe(port, router)
}
