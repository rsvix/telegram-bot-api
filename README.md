# Telegram Bot API

A small HTTP + Telegram bot service that receives POST requests to send messages to Telegram chats. Built with Go, Echo (HTTP), and www.github.com/go-telegram/bot.

---

## Features

- Send messages to Telegram via POST /send-message
- Read bot token from environment variable or file:
  - `TELEGRAM_BOT_TOKEN_FILE` (path to file containing token)
  - `TELEGRAM_BOT_TOKEN` (token value)
- Optional disable of link previews per message

---

Server listens on `:1323` by default.

---

## API

### POST /send-message

- Content-Type: `application/json`
- Body (JSON):

```json
{
  "chat_id": 123456789,
  "message": "Hello world",
  "disable_link_preview": true
}
```

Notes:
- `chat_id` is an integer (`int64`).
- `message` is required (non-empty).
- `disable_link_preview` is optional; when omitted it defaults to `true`.

Responses:
- `200` — `{"status":"message sent"}` (success)
- `400` — `{"error":"..."}` (client error, invalid JSON or missing fields)
- `500` — `{"error":"..."}` (server error, e.g., failed to send)

---

## Configuration & Environment

- `TELEGRAM_BOT_TOKEN_FILE` — path to file with bot token (preferred for secrets)
- `TELEGRAM_BOT_TOKEN` — raw bot token (fallback)

---

## License

MIT
