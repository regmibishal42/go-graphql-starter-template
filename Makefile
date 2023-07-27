init:
	go run github.com/99designs/gqlgen init
generate:
	go run github.com/99designs/gqlgen generate
serve:
	go run cmd/server.go
test:
	go run test.go