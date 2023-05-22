package handlers

import (
	"github.com/MichaelYoung87/kundbild-public/application/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type PlanetsHandler struct {
	service *services.PlanetsService
}

func NewPlanetsHandler(service *services.PlanetsService) *PlanetsHandler {
	return &PlanetsHandler{
		service: service,
	}
}

func (h *PlanetsHandler) HandlePlanetsURLEnding(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse id",
		})
	}

	// Calls the service method with this id.
	planets, err := h.service.GetPlanetsFromAPIClientByURLEnding(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching planets data",
		})
	}

	// Return the response with the planets data
	return c.JSON(planets)
}
