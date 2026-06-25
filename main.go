package main

import (
	"os"

	"github.com/abdimk/vercelgo/internals"
)

func main() {
	
	// Vercel injects PORT; fall back to 3000 for local dev
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	Server := internals.NewServer(0)
	Server.Run(port)

	
}
