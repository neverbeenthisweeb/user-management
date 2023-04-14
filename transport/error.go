package transport

import (
	"encoding/json"
	"net/http"
	"usermanagement/service"

	"github.com/gofiber/fiber/v2"
)

const (
	errMsgMissingOrMalformedJWT = "Missing or malformed JWT"
)

type errResponse struct {
	ErrorMessage string `json:"error_message"`
}

func ErrHandler(c *fiber.Ctx, err error) error {
	code := http.StatusInternalServerError

	if v, ok := err.(service.ServiceErr); ok {
		code = v.Code
	}

	if err.Error() == errMsgMissingOrMalformedJWT {
		code = http.StatusUnauthorized
	}

	// Send error response
	errResp := errResponse{
		ErrorMessage: err.Error(),
	}

	b, errMarshal := json.Marshal(errResp)
	if errMarshal != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error: " + errMarshal.Error())
	}

	errSend := c.Status(code).Send(b)
	if errSend != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error: " + errSend.Error())
	}

	return nil
}
