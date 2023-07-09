package main

import (
	"fmt"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/client/bidirectional_streaming"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/client/unary"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/rest"
	"github.com/Inspirate789/golang-sandbox/internal/models"

	_ "github.com/Inspirate789/golang-sandbox/swagger"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func calculationCallback(calculators ...models.CalculatorExternal) models.CalculationCallback {
	return func(input models.Calculation) (models.Calculation, error) {
		var err error
		for _, calculator := range calculators {
			input, err = calculator.Calculate(input)
			if err != nil {
				return models.Calculation{}, err
			}
		}

		return models.DefaultCalculation(input)
	}
}

func setupWebApp(port string, callback models.CalculationCallback, logOutput io.Writer) *fiber.App {
	app := fiber.New()

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          fmt.Sprintf("http://localhost:%s/swagger/doc.json", port),
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
		Output: logOutput,
	}))

	api := app.Group("/api/v1")
	rest.SetupDelivery(api, logOutput, callback)

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

// @title			Sandbox API
// @version			0.1.0
// @description		This is a sandbox API.
// @contact.name	API Support
// @contact.email	andreysapozhkov535@gmail.com
// @host			localhost:8080
// @BasePath		/api/v1
// @Schemes			http
func main() {
	mainLogger := log.New(os.Stdout, "Main: ", log.LstdFlags)

	uc := unary.NewClient()
	err := uc.Open("localhost:5300")
	if err != nil {
		mainLogger.Fatal(err)
	}

	bc := bidirectional_streaming.NewClient()
	err = bc.Open("localhost:5301")
	if err != nil {
		mainLogger.Fatal(err)
	}

	app := setupWebApp("8080", calculationCallback(uc, bc), os.Stdout)

	go func() {
		err = app.Listen(":8080") // TODO: add TLS config
		if err != nil {
			mainLogger.Fatal(err)
		}
	}()
	mainLogger.Println("Server started at port 8080")

	teardownWebApp(mainLogger, app)
}
