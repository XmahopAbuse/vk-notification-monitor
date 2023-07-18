package sender

import "github.com/mymmrac/telego"

type TelegramSender struct {
	bot telego.Bot
}

func (t *TelegramSender) Send() error {
	return nil
}
