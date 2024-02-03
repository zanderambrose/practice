scraper:
	go run cmd/scraper/main.go;

api:
	go run github.com/cosmtrek/air

node:
	cd ./node-api && npm run dev

db:
	docker compose up -d db;

down:
	docker compose down;
