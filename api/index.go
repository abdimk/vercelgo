package handler

import (
	"net/http"
	"time"

	"github.com/abdimk/vercelgo/schema"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

// Handler is the entry point Vercel calls for serverless functions.
func Handler(w http.ResponseWriter, r *http.Request) {
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

	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(emotion)
	})

	// Ensure Fiber receives the correct request URI for routing
	r.RequestURI = r.URL.String()
	adaptor.FiberApp(app).ServeHTTP(w, r)
}
