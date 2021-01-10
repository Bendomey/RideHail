run:
	reflex -r '\.go' -s -- sh -c "go run cmd/server.go"


# for generating grapqhl
generate-graph:
	go run github.com/99designs/gqlgen generate