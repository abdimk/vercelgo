package main

import (
	"log"
	"os"
	"time"

	"github.com/abdimk/vercelgo/schema"
	"github.com/gofiber/fiber/v3"
)

func main() {
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

	// Vercel injects PORT; fall back to 3000 for local dev
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(emotion)
	})

	log.Fatal(app.Listen(":" + port))
}
