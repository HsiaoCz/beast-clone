package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/twitter/app"
	"github.com/HsiaoCz/beast-clone/twitter/db"
	"github.com/HsiaoCz/beast-clone/twitter/etc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if err := etc.ParseConfig(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(etc.Conf.App.MongoUri))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			log.Fatal(err)
		}
	}()

	var (
		logger           = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
		port             = etc.Conf.App.Port
		userColl         = client.Database(etc.Conf.App.DBName).Collection(etc.Conf.App.UserColl)
		postColl         = client.Database(etc.Conf.App.DBName).Collection(etc.Conf.App.PostColl)
		commentColl      = client.Database(etc.Conf.App.DBName).Collection(etc.Conf.App.CommentColl)
		mongoUserCase    = db.NewMongoUserCase(client, userColl)
		mongoPostCase    = db.NewMongoPostStore(client, postColl)
		mongoCommentCase = db.NewMongoCommentStore(client, commentColl)
		dbs              = &db.DBS{Uc: mongoUserCase, Pc: mongoPostCase, CS: mongoCommentCase}
		userApp          = app.NewUserApp(dbs)
		postApp          = app.NewPostApp(dbs)
		commentApp       = app.NewCommentApp(dbs)
		router           = chi.NewMux()
	)

	slog.SetDefault(logger)

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
	router.Delete("/posts/{pid}", app.TransferHandlerfunc(postApp.HandleDeletePost))

	// TODO need group router comment
	// posts handlers need auth middleware
	router.Post("/comments", app.TransferHandlerfunc(commentApp.HandleCreateComment))
	router.Delete("/comments/{cid}", app.TransferHandlerfunc(commentApp.HandleDeleteCommentByID))

	slog.Info("the server is running", "listen address", port)
	http.ListenAndServe(port, router)
}
