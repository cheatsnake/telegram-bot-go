build:
	go build cmd/main.go
docker-build:
	docker build . -t telegram-bot-go
docker-start:
	docker run -e BOT_TOKEN=$(BOT_TOKEN) -d --rm --name telegram-bot-go telegram-bot-go
docker-stop:
	docker stop telegram-bot-go