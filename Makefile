run:
	reflex -r '\.go' -s -- sh -c "go run cmd/server.go"

run-accounts:
	reflex -r '\.go' -s -- sh -c "go run account/cmd/main.go"

# for generating grapqhl
generate-graph:
	go run github.com/99designs/gqlgen generate