.PHONY:
.SILENT:

build:
	docker-compose exec app go build -o ./.bin/bot cmd/bot/main.go

run: build
	docker-compose exec app ./.bin/bot
