package app

import (
	"context"
	"log"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"mongorestapi3/internal/handler"
	"mongorestapi3/internal/service"
)

type App struct {
	Router         *mux.Router
	StudentHandler *handler.StudentHandler
}

func (a *App) Initialize(ctx context.Context, conf *Config) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.URI))
	if err != nil {
		log.Fatal("Could not connect to database !!!")
	}
	db := client.Database(conf.Mongo.Database)

	studentService := service.NewStudentService(db)
	a.StudentHandler = handler.NewStudentHandler(studentService)

	a.Router = mux.NewRouter()
	a.setRouters()
}
