package common

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// IsAdmin 4Head
func IsAdmin(user *tgbotapi.User) bool {
	var admins = map[int]bool{
		71524437:  true,
		105513756: true,
	}
	if _, exists := admins[user.ID]; exists {
		return true
	}
	return false
}
