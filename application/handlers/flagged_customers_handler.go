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

type FlaggedCustomersHandler struct {
	flaggedCustomersService *services.FlaggedCustomersService
}
type flaggedCustomersPostPayload struct {
	People  people.People   `json:"people"`
	Planets planets.Planets `json:"planets"`
}

func NewFlaggedCustomersHandler(flaggedCustomersService *services.FlaggedCustomersService) *FlaggedCustomersHandler {
	return &FlaggedCustomersHandler{
		flaggedCustomersService: flaggedCustomersService,
	}
}

func (fch *FlaggedCustomersHandler) PostFlaggedCustomers(c *fiber.Ctx) error {
	var payload flaggedCustomersPostPayload
	err := json.Unmarshal(c.Body(), &payload)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	if !fch.flaggedCustomersService.CheckMatch(&payload.People, &payload.Planets) {
		return c.Status(http.StatusBadRequest).JSON("Provided data does not match the hardcoded combinations.")
	}

	flaggedCustomers, isNew, err := fch.flaggedCustomersService.SaveFlaggedCustomersToDatabase(&payload.People, &payload.Planets)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	if isNew {
		return c.Status(http.StatusOK).JSON("Successfully saved flagged customers with ID: " + strconv.Itoa(int(flaggedCustomers.FlaggedCustomersID)))
	} else {
		return c.Status(http.StatusOK).JSON("Flagged customers with ID: " + strconv.Itoa(int(flaggedCustomers.FlaggedCustomersID)) + " already exists and will not be added to the database.")
	}
}
