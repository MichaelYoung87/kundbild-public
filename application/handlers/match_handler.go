package handlers

import (
	"github.com/MichaelYoung87/kundbild-public/application/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type MatchHandler struct {
	peopleService           *services.PeopleService
	planetsService          *services.PlanetsService
	flaggedCustomersService *services.FlaggedCustomersService
}

func NewMatchHandler(peopleService *services.PeopleService, planetsService *services.PlanetsService, flaggedCustomersService *services.FlaggedCustomersService) *MatchHandler {
	return &MatchHandler{
		peopleService:           peopleService,
		planetsService:          planetsService,
		flaggedCustomersService: flaggedCustomersService,
	}
}

func (h *MatchHandler) CheckMatch(c *fiber.Ctx) error {
	// Parse the URL endings from the request
	peopleURLEnding, err := strconv.Atoi(c.Params("peopleURLEnding"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid People URL ending")
	}
	planetsURLEnding, err := strconv.Atoi(c.Params("planetsURLEnding"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Planets URL ending")
	}

	// Fetch the People and Planets entities using their URL endings
	people, err := h.peopleService.GetPeopleFromAPIClientByURLEnding(peopleURLEnding)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("People not found")
	}
	planets, err := h.planetsService.GetPlanetsFromAPIClientByURLEnding(planetsURLEnding)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Planets not found")
	}

	// Checks if the names of the fetched People and Planets match any of the combinations in the matchingCombinations map inside flagged_customers_service.go
	match := h.flaggedCustomersService.CheckMatch(people, planets)

	// Send the result as a JSON response
	return c.JSON(fiber.Map{"match": match})
}
