package jokenpo

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type HTTPHandler struct {
	service *Service
}

func NewHTTPHandler(app *fiber.App, s *Service) {
	h := &HTTPHandler{
		service: s,
	}

	app.Post("/jokenpo/:option", h.Play)
}

func (h *HTTPHandler) Play(c fiber.Ctx) error {
	playerOption, err := strconv.Atoi(c.Params("option"))
	if err != nil {
		return err
	}

	if playerOption < 0 || playerOption > 2 {
		return c.Status(http.StatusBadRequest).
			JSON(map[string]string{
				"message": "valid options are: 0 (Pedra), 1 (Papel), or 2 (Tesoura)",
			})
	}

	computerOption := rand.Intn(2)

	winner, options := h.service.Play(playerOption, computerOption)

	return c.JSON(map[string]string{
		"user":     options[0].String(),
		"computer": options[1].String(),
		"winner":   winner.String(),
	})
}
