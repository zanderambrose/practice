scraper:
	go run cmd/scraper/main.go;

api:
	go run github.com/cosmtrek/air

up:
	docker compose up -d --force-recreate;

db:
	docker compose up -d db;

db-shell:
	docker exec -it whoshittin-db-1 bash -c "mongosh"

down:
	docker compose down;
