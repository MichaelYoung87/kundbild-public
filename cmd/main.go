package main

import (
	"github.com/MichaelYoung87/kundbild-public/application/handlers"
	"github.com/MichaelYoung87/kundbild-public/application/services"
	"github.com/MichaelYoung87/kundbild-public/infrastructure/apiclients/peopleapi"
	"github.com/MichaelYoung87/kundbild-public/infrastructure/apiclients/planetsapi"
	"github.com/MichaelYoung87/kundbild-public/infrastructure/persistence/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func GetSWAPI() string {
	SWAPI := os.Getenv("SWAPI_URL")
	return SWAPI
}

func main() {
	USERNAME := os.Getenv("DB_USERNAME")
	PASSWORD := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DATABASENAME := os.Getenv("DB_NAME")

	db, err := database.ConnectToDB(USERNAME, PASSWORD, HOST, PORT, DATABASENAME)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	SWAPI := GetSWAPI()
	httpClient := &http.Client{}

	// Initialize repositories
	peopleRepo := database.NewPeopleDBRepository(db)
	planetsRepo := database.NewPlanetsDBRepository(db)
	flaggedCustomersRepo := database.NewFlaggedCustomersDBRepository(db)
	linkedCustomersRepo := database.NewLinkedCustomersDBRepository(db)

	// Initialize API clients
	peopleAPI := peopleapi.NewPeopleAPIClient(SWAPI, httpClient)
	planetsAPI := planetsapi.NewPlanetsAPIClient(SWAPI, httpClient)

	// Initialize services
	peopleService := services.NewPeopleService(peopleRepo, peopleAPI)
	planetsService := services.NewPlanetsService(planetsRepo, planetsAPI)
	flaggedCustomersService := services.NewFlaggedCustomersService(flaggedCustomersRepo, peopleRepo, planetsRepo)
	linkedCustomersService := services.NewLinkedCustomersService(linkedCustomersRepo, peopleRepo, planetsRepo)

	// Initialize handlers
	flaggedCustomersHandler := handlers.NewFlaggedCustomersHandler(flaggedCustomersService)
	linkedCustomersHandler := handlers.NewLinkedCustomersHandler(linkedCustomersService)
	matchHandler := handlers.NewMatchHandler(peopleService, planetsService, flaggedCustomersService)
	peopleHandler := handlers.NewPeopleHandler(peopleService)
	planetsHandler := handlers.NewPlanetsHandler(planetsService)

	// Create new Fiber instance
	app := fiber.New()

	// Enable CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	// Define application routes.
	app.Get("/people/:id", peopleHandler.HandlePeopleURLEnding)
	app.Get("/planets/:id", planetsHandler.HandlePlanetsURLEnding)
	app.Get("/match/:peopleURLEnding/:planetsURLEnding", matchHandler.CheckMatch)
	app.Post("/flagged_customers", flaggedCustomersHandler.PostFlaggedCustomers)
	app.Post("/linked_customers", linkedCustomersHandler.PostLinkedCustomers)

	// Start the Fiber server
	log.Fatal(app.Listen(":8000"))
}

//	"Luke Skywalker"	= people/1/		"Hoth"		= planets/4/ - Hard coded in flagged_customers_service.go and will be a match for flagging
//	"Leia Organa"		= people/5/		"Tatooine"	= planets/1/ - Hard coded in flagged_customers_service.go and will be a match for flagging
// 	"Darth Vader"		= people/4/		"Dagobah"	= planets/5/ - Hard coded in flagged_customers_service.go and will be a match for flagging
// 	"C-3PO"				= people/2/		"Yavin IV"	= planets/3/ - Hard coded in flagged_customers_service.go and will be a match for flagging
//	"R2-D2"				= people/3/		"Alderaan"	= planets/2/ - Hard coded in flagged_customers_service.go and will be a match for flagging
//	"Owen Lars"			= people/6/		"Bespin"	= planets/6/ - Hard coded in flagged_customers_service.go and will be a match for flagging
