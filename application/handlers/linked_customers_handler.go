package handlers

import (
	"encoding/json"
	"github.com/MichaelYoung87/kundbild-public/application/services"
	"github.com/MichaelYoung87/kundbild-public/domain/people"
	"github.com/MichaelYoung87/kundbild-public/domain/planets"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type LinkedCustomersHandler struct {
	linkedCustomersService *services.LinkedCustomersService
}
type linkedCustomersPostPayload struct {
	People  people.People   `json:"people"`
	Planets planets.Planets `json:"planets"`
}

func NewLinkedCustomersHandler(linkedCustomersService *services.LinkedCustomersService) *LinkedCustomersHandler {
	return &LinkedCustomersHandler{
		linkedCustomersService: linkedCustomersService,
	}
}

func (fch *LinkedCustomersHandler) PostLinkedCustomers(c *fiber.Ctx) error {
	var payload linkedCustomersPostPayload
	err := json.Unmarshal(c.Body(), &payload)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	linkedCustomers, isNew, err := fch.linkedCustomersService.SaveLinkedCustomersToDatabase(&payload.People, &payload.Planets)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	if isNew {
		return c.Status(http.StatusOK).JSON("Successfully saved linked customers with ID: " + strconv.Itoa(int(linkedCustomers.LinkedCustomersID)))
	} else {
		return c.Status(http.StatusOK).JSON("Linked customers with ID: " + strconv.Itoa(int(linkedCustomers.LinkedCustomersID)) + " already exists and will not be added to the database.")
	}
}
