up:
	docker compose up --build

down:
	docker compose down

logs:
	docker compose logs -f

test:
	go test ./...

build:
	docker build -t identity-api:local .

clean:
	docker compose down -v