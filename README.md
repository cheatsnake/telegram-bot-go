# Telegram Bot Go

A minimal implementation of a telegram bot for receiving and sending messages in pure Go. Use this as an initial template for your cool bots.

## Usage

1. Clone this repo

```sh
git clone https://github.com/cheatsnake/telegram-bot-go.git
```

```sh
cd ./telegram-bot-go
```

2. Build application binary

```sh
make build
```

3. Run it

```sh
./main "paste_your_token_here"
```

## 🐳 Docker container startup

1. Start building proccess:

```sh
make docker-build
```

2. Running a container with the specified token:

```sh
make docker-start BOT_TOKEN="paste_your_token_here"
```

To stop the container and delete it use:

```sh
make docker-stop
```
