package main

import (
	"fmt"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/interfaces"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/server/bidirectional_streaming"
	"github.com/Inspirate789/golang-sandbox/internal/models"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func teardownServer(mainLogger *log.Logger, server interfaces.Server) {
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	mainLogger.Println("Shutdown server ...")

	err := server.Close()
	if err != nil {
		mainLogger.Fatal(fmt.Sprintf("Server shutdown: %v", err))
	}
	mainLogger.Println("Server exited")
}

func main() {
	mainLogger := log.New(os.Stdout, "Main: ", log.LstdFlags)

	server, err := bidirectional_streaming.NewServer(":5301", models.DefaultCalculation)
	if err != nil {
		mainLogger.Fatal(err)
	}

	go func() {
		err = server.Serve()
		if err != nil {
			mainLogger.Fatal(err)
		}
	}()
	mainLogger.Println("Server started at port 5301")

	teardownServer(mainLogger, server)
}
