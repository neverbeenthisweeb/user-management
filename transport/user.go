package transport

import (
	"net/http"
	"usermanagement/model"
	"usermanagement/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog/log"
)

const (
	defaultPort = ":4040"
)

type controller struct {
	svc service.User
}

func NewController(svc service.User) *controller {
	return &controller{
		svc: svc,
	}
}

func (t *controller) Login(c *fiber.Ctx) error {
	logger := log.With().
		Str("method", "controller.Login").
		Str("requestid", c.Context().Value("requestid").(string)).
		Logger()

	req := model.User{}

	err := c.BodyParser(&req)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse body")
		return err
	}

	token, err := t.svc.Login(c.Context(), req)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to login")
		return err
	}

	return c.Status(http.StatusOK).JSON(token)
}

func (t *controller) Start(app *fiber.App) {
	// Middleware
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${locals:requestid} ${status} - ${latency} ${method} ${path}​\n",
	}))

	// Route
	app.Post("/login", t.Login)

	// Listen
	err := app.Listen(defaultPort)
	if err != nil {
		panic(err)
	}
}
