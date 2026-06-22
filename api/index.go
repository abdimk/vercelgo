package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abdimk/vercelgo/schema"
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emotion)
}
