package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cheatsnake/telegram-bot-go/internal/telegram"
)

const defaultMsgLimit = 10 // How many messages the bot has to process in one polling

func main() {
	bot := telegram.New(mustToken())
	currentOffset := 1

	// Check is bot token valid
	botInfo, err := bot.GetMe()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Bot @%s is running...\n", botInfo.Username)

	// Running periodic pollings
	for {
		updates, err := bot.Updates(currentOffset, defaultMsgLimit)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Need to get only fresh updates
		if len(updates) > 0 {
			currentOffset = updates[len(updates)-1].ID + 1
		}

		// Process each message we receive
		for i := range updates {
			// Just send the user his own message
			go bot.SendMessage(updates[i].Message.Chat.ID, updates[i].Message.Text)
		}

		// Polling frequency
		time.Sleep(2 * time.Second)
	}
}

func mustToken() string {
	token := os.Getenv("BOT_TOKEN")
	if len(token) > 0 {
		return token
	}

	if len(os.Args) < 2 {
		log.Fatal("telegram bot API token is required")
	}

	return os.Args[1]
}
