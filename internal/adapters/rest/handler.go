package rest

import (
	"github.com/Inspirate789/golang-sandbox/internal/models"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
)

type delivery struct {
	logger   *log.Logger
	callback models.CalculationCallback
}

func newDelivery(callback models.CalculationCallback, logOutput io.Writer) *delivery {
	return &delivery{
		logger:   log.New(logOutput, "Delivery: ", log.LstdFlags),
		callback: callback,
	}
}

func SetupDelivery(api fiber.Router, logOutput io.Writer, callback models.CalculationCallback) {
	d := newDelivery(callback, logOutput)
	api.Get("/call", d.call)
}

// call godoc
//
//	@Summary		Call the REST API.
//	@Description	call the REST API
//	@Tags			Control
//	@Param			input	query	int64	true	"Calculation input"
//	@Produce		json
//	@Success		200	{object}	models.Calculation
//	@Failure		422	{object}	fiber.Map
//	@Failure		500	{object}	fiber.Map
//	@Router			/call [get]
func (d *delivery) call(ctx *fiber.Ctx) error {
	input := ctx.QueryInt("input", -1)
	d.logger.Printf("Received calculation input: %d", input)
	if input < 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "Incorrect calculation input (must be non-negative integer number)",
		})
	}
	res := models.Calculation{
		Base:   uint64(input),
		Result: 1,
	}

	res, err := d.callback(res)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}
