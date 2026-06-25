package internals

import (
	"log"
	"time"

	"github.com/abdimk/vercelgo/schema"
	"github.com/gofiber/fiber/v3"
)



type Server struct{
	ListingAdder any
	db *string
}

func NewServer(listingAdder any)*Server{
	return &Server{ListingAdder: listingAdder, db: nil}
}
func (s *Server) Run(port string){
	emotion := schema.Emotions{
		Name:      "Lost",
		Intensity: 0.8,
		Valence:   1.0,
		Duration:  30 * time.Minute,
		Trigger:   "Working on the payment API",
		StartTime: time.Now(),
		Target:    "Myself",
		PhysicalEffects: []string{
			"Smiling",
			"Relaxed muscles",
		},
		Thoughts: []string{
			"I did a good job",
			"This was worth the effort",
		},
		Expression: "Smile",
	}
	
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(emotion)
	})
	
	
	app.Post("/user",func(c fiber.Ctx) error{
		var user schema.User;
		
		if err := c.Bind().Body(&user); err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":err.Error(),
			})
		}
		return c.JSON(user)
	})
	
	log.Fatal(app.Listen(":" + port))
}
