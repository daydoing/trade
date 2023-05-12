package context

import (
	"gopkg.in/tucnak/telebot.v2"
)

type NotifyBot struct {
	*telebot.Bot
	uid int
}

func initNotifyBot(token string, uid int) (*NotifyBot, error) {
	n, err := telebot.NewBot(telebot.Settings{Token: token})
	if err != nil {
		return nil, err
	}

	return &NotifyBot{n, uid}, nil
}

func (n *NotifyBot) Notify(content string) error {
	_, err := n.Send(&telebot.Chat{ID: int64(n.uid)}, content)
	return err
}
