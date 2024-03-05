.PHONY: build

build:
	docker compose build my_app_app

up:
	docker compose up --build -d my_app_dev

down:
	docker compose down

air:
	docker compose exec air bash