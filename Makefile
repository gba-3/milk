run:
	docker-compose exec appserver go run main.go
restart:
	docker-compose down && docker-compose up -d --build
