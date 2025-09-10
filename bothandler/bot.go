package bothandler

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func NewTelegramBot(ctx context.Context, token string) (*bot.Bot, error) {
	if token == "" {
		return nil, fmt.Errorf("token required")
	}

	opts := []bot.Option{bot.WithDefaultHandler(handler)}

	b, err := bot.New(token, opts...)
	if err != nil {
		return nil, err
	}

	go b.Start(ctx)
	return b, nil
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update == nil || update.Message == nil {
		return
	}

	message := update.Message.Text
	chatId := update.Message.Chat.ID
	log.Println("received message:", message)
	log.Println("received from id:", chatId)

	var err error
	switch message {
	case "/hello":
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatId,
			Text:   "world",
		})
	case "/chat-id":
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatId,
			Text:   fmt.Sprintf("This chat ID: %v", chatId),
		})
	default:
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatId,
			Text:   "Hi, i'm a telegram Bot",
		})
	}

	if err != nil {
		log.Println("SendMessage error:", err)
	}
}
