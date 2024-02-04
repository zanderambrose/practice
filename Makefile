scraper:
	go run cmd/scraper/main.go;

api:
	go run github.com/cosmtrek/air

up:
	docker compose up -d --force-recreate -V;

dev:
	docker compose up -d --build --force-recreate -V;

db:
	docker compose up -d db;

db-shell:
	docker exec -it whoshittin-db-1 bash -c "mongosh mongodb://root:examplepassword@db:27017"

down:
	docker compose down;
