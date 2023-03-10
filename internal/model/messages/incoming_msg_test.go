package messages

import (
	mocks "github.com/Gamilkarr/route256-telegram-bot/internal/mocks/messages"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)
	sender.EXPECT().SendMessage("Привет", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: 123,
	})
	assert.NoError(t, err)
}

func TestOnUnknownCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)
	sender.EXPECT().SendMessage("Прости, я не знаю такой команды", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "/some_comand",
		UserID: 123,
	})
	assert.NoError(t, err)
}
