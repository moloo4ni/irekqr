package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/skip2/go-qrcode"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is not set")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() && update.Message.Text == "" {
			continue
		}

		if update.Message.IsCommand() && update.Message.Command() == "start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Send me any text or URL and I'll generate a QR code for it.")
			bot.Send(msg)
			continue
		}

		text := strings.TrimSpace(update.Message.Text)
		if text == "" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Text can't be empty.")
			bot.Send(msg)
			continue
		}

		if len(text) > 4000 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Text is too long (max 4000 characters).")
			bot.Send(msg)
			continue
		}

		png, err := qrcode.Encode(text, qrcode.Medium, 256)
		if err != nil {
			log.Printf("Error generating QR: %v", err)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Something went wrong. Try again.")
			bot.Send(msg)
			continue
		}

		photo := tgbotapi.FileBytes{
			Name:  "qr.png",
			Bytes: png,
		}
		msg := tgbotapi.NewPhoto(update.Message.Chat.ID, photo)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			log.Printf("Error sending photo: %v", err)
		}
	}
}
