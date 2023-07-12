package main

import (
	"fmt"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/websocket/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func setupWebApp(logOutput io.Writer) *fiber.App {
	engine := html.New("./internal/adapters/websocket/client/logger", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
		Output: logOutput,
	}))

	api := app.Group("/api/v1")
	server.SetupDelivery(api, logOutput)

	return app
}

func teardownWebApp(mainLogger *log.Logger, app *fiber.App) {
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	mainLogger.Println("Shutdown app ...")

	err := app.Shutdown()
	if err != nil {
		mainLogger.Fatal(fmt.Sprintf("App shutdown: %v", err))
	}
	mainLogger.Println("App exited")
}

func main() {
	mainLogger := log.New(os.Stdout, "Main: ", log.LstdFlags)

	app := setupWebApp(os.Stdout)

	go func() {
		err := app.Listen(":30081") // TODO: add TLS config
		if err != nil {
			mainLogger.Fatal(err)
		}
	}()
	mainLogger.Println("Server started at port 30081")

	teardownWebApp(mainLogger, app)
}
