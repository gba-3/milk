run:
	docker-compose exec appserver go run main.go
restart:
	docker-compose down && docker-compose up -d --build
lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.35.2 golangci-lint run -v
