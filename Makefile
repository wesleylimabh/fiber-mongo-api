all: run

run:
	@go run main.go

dev: 
	@air

doc:
	@swag init -g routes/swagger.go

build:
	@go build -o bin/app main.go
