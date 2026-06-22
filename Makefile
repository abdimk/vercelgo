build:
	@go build -o bin/main main.go

run:build
	@bin/main

vercel:
	@echo "Deploy to Vercel: vercel --prod"

all: build

