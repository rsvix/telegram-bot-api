package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/telegram-bot-api/bothandler"
	"github.com/rsvix/telegram-bot-api/sendmessage"
)

func loadBotToken() (string, error) {
	if path, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN_FILE"); ok {
		b, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("error reading TELEGRAM_BOT_TOKEN_FILE: %w", err)
		}
		return strings.TrimSpace(string(b)), nil
	}
	if v, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN"); ok {
		return v, nil
	}
	return "", fmt.Errorf("TELEGRAM_BOT_TOKEN_FILE or TELEGRAM_BOT_TOKEN must be set")
}

func main() {
	token, err := loadBotToken()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bothandler.NewTelegramBot(ctx, token)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.POST("/send-message", sendmessage.SendMessageHandler(b).Serve)
	e.Logger.Fatal(e.Start(":1323"))
}
