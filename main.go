package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"coginfra/apps/api"
	"coginfra/apps/cli"
	"coginfra/configs"
	"coginfra/storage"
	"coginfra/utils"
)

func main() {
	isDebugMode := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	c := configs.LoadConfig(*isDebugMode)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a WaitGroup to synchronize server start and menu display

	var store storage.Storage = storage.ConnectToMongo(c.MongoDSN)
	// var store storage.Storage = &storage.MemoryStorage{}

	s := api.NewServer(":8080", store)

	api.InitServer(s)

	for !s.HasStarted {
	}

	cli.StartMenu(store)

	// Wait for termination signal
	<-signalChan

	utils.Logger.Println("Shutting down gracefully...")
	store.Disconnect()
	// Todo: Add cleanup code if needed

	utils.Logger.Println("Exiting.")
}
