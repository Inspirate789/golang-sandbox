package server

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
)

type delivery struct {
	logger *log.Logger
}

func newDelivery(logOutput io.Writer) *delivery {
	return &delivery{
		logger: log.New(logOutput, "Delivery: ", log.LstdFlags),
	}
}

func SetupDelivery(api fiber.Router, logOutput io.Writer) {
	d := newDelivery(logOutput)
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Test",
		})
	})
	api.Use("/ws", d.wsHandler)
	api.Get("/ws/:id", websocket.New(d.wsConnHandler))
}

func (d *delivery) wsHandler(ctx *fiber.Ctx) error {
	// IsWebSocketUpgrade returns true if the client
	// requested upgrade to the WebSocket protocol.
	if websocket.IsWebSocketUpgrade(ctx) {
		ctx.Locals("allowed", true)
		return ctx.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (d *delivery) wsConnHandler(ctx *websocket.Conn) {
	// ctx.Locals is added to the *websocket.Conn
	d.logger.Println(ctx.Locals("allowed"))  // true
	d.logger.Println(ctx.Params("id"))       // 123
	d.logger.Println(ctx.Query("v"))         // 1.0
	d.logger.Println(ctx.Cookies("session")) // ""

	var (
		msgType int
		msg     []byte
		err     error
	)
	for {
		if msgType, msg, err = ctx.ReadMessage(); err != nil {
			d.logger.Println("read:", err)
			break
		}
		d.logger.Printf("receive (type %d): %s", msgType, msg)

		if err = ctx.WriteJSON(fiber.Map{"msg": string(msg)}); err != nil {
			d.logger.Println("write:", err)
			break
		}
	}
}
