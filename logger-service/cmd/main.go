package main

import (
	"github.com/Noah-Wilderom/grpc-server/logger-service/database"
	"github.com/Noah-Wilderom/grpc-server/logger-service/database/models"
	"github.com/Noah-Wilderom/grpc-server/logger-service/handlers"
	"os"
	"sync"
)

func main() {
	ctx, client := database.ConnectToDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db := client.Database(os.Getenv("MONGO_DATABASE"))

	logModel := models.NewLogModel(db)
	logHandler := handlers.NewLogHandler(logModel)

	app := NewServer(logHandler)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		app.Run()
	}()

	wg.Wait()
}
