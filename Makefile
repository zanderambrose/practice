scraper:
	go run cmd/scraper/main.go;

api:
	go run github.com/cosmtrek/air

db:
	docker compose up -d db;
