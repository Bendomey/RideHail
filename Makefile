run:
	reflex -r '\.go' -s -- sh -c "go run cmd/server.go"

start:
	heroku local web

# for generating grapqhl
generate-graph:
	go run github.com/99designs/gqlgen generate