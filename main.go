package main

import (
	"log"
	"time"

	"github.com/abdimk/vercelgo/schema"
	"github.com/gofiber/fiber/v3"
)

func main(){
	emotion := schema.Emotions{
    Name:      "Happiness",
    Intensity: 0.8,
    Valence:   1.0,
    Duration:  30 * time.Minute,
    Trigger:   "Finished a project",
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
	// fiber golang
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx)error{
		return c.JSON(emotion)
		
	})
	
	log.Fatal(app.Listen(":3000"))
	
}