run:
	@templ generate
	@go run cmd/api/main.go

build:
	@echo "Building..."
	@npm run build:prod
	@templ generate
	@go build -o ./tmp/main cmd/api
