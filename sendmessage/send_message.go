package sendmessage

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/labstack/echo/v4"
)

type sendMessageReq struct {
	ChatID             int64  `json:"chat_id"`
	Message            string `json:"message"`
	DisableLinkPreview *bool  `json:"disable_link_preview,omitempty"`
}

type sendMessageHandlerParams struct {
	telegramBot *bot.Bot
}

func SendMessageHandler(b *bot.Bot) *sendMessageHandlerParams {
	return &sendMessageHandlerParams{
		telegramBot: b,
	}
}

func (h *sendMessageHandlerParams) Serve(c echo.Context) error {
	var req sendMessageReq
	// if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
	// }
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
	}

	req.Message = strings.TrimSpace(req.Message)
	if req.Message == "" || req.ChatID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing chat_id or message"})
	}

	disablePreview := true // default
	if req.DisableLinkPreview != nil {
		disablePreview = *req.DisableLinkPreview
	}

	log.Printf("\n# --------------------\nchat_id: %v\nmessage: %s\ndisableLinkPreview: %v\n", req.ChatID, req.Message, disablePreview)

	if _, err := h.telegramBot.SendMessage(c.Request().Context(), &bot.SendMessageParams{
		ChatID: req.ChatID,
		Text:   req.Message,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: &disablePreview,
		},
		ParseMode: "HTML",
	}); err != nil {
		log.Println("SendMessage error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "couldn't send message"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "message sent"})
}
