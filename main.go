package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	"huydevbot/Message"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	botToken := os.Getenv("TOKEN")
	url := os.Getenv("URL_WEBHOOK")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set up a webhook on Telegram side
	_ = bot.SetWebhook(&telego.SetWebhookParams{
		URL: url + "/bot" + bot.Token(),
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo()
	fmt.Printf("Webhook Info: %+v\n", info)

	// Get an update channel from webhook.
	// (more on configuration in examples/updates_webhook/main.go)
	updates, _ := bot.UpdatesViaWebhook("/bot" + bot.Token())

	// Start server for receiving requests from the Telegram
	go func() {
		_ = bot.StartWebhook("0.0.0.0:8080")
	}()

	// Stop reviving updates from update channel and shutdown webhook server
	defer func() {
		_ = bot.StopWebhook()
	}()

	// Loop through all updates when they came
	for update := range updates {
		Message.HandleMessage(update, bot)
	}
}
