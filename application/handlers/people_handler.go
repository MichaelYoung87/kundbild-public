package handlers

import (
	"github.com/MichaelYoung87/kundbild-public/application/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type PeopleHandler struct {
	service *services.PeopleService
}

func NewPeopleHandler(service *services.PeopleService) *PeopleHandler {
	return &PeopleHandler{
		service: service,
	}
}

func (h *PeopleHandler) HandlePeopleURLEnding(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse id",
		})
	}

	// Calls the service method with this id.
	people, err := h.service.GetPeopleFromAPIClientByURLEnding(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching people data",
		})
	}

	// Return the response with the people data
	return c.JSON(people)
}
