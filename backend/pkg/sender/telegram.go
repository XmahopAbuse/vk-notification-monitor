package sender

import (
	"github.com/mymmrac/telego"
	"log"
)

type TelegramSender struct {
	bot *telego.Bot
}

func NewTelegramSender(bot *telego.Bot) *TelegramSender {
	return &TelegramSender{bot: bot}
}

func (t *TelegramSender) Send() error {
	log.Printf("send to telegram")
	return nil
}
