package helpers

import (
	"github.com/nicklaw5/helix/v2"
	"madnessBot/common/logger"
	"madnessBot/config"
)

// GetTwitchUser get user by login
func GetTwitchUser(login string) (*helix.User, error) {
	resp, err := config.Config.Twitch.Client().GetUsers(&helix.UsersParams{
		Logins: []string{login},
	})

	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to get twitch user")
		return nil, err
	}

	return &resp.Data.Users[0], nil
}

//GetTwitchUserIDByLogin get userID by Twitch login
func GetTwitchUserIDByLogin(login string) (string, bool) {
	user, err := GetTwitchUser(login)

	if err != nil {
		logger.Log.Error().Err(err).Msg("Request failed")
		return "", false
	}

	return user.ID, user.ID != ""
}

//SendEventSubMessage sends a message to the Twitch Hub
func SendEventSubMessage(channel string, eventType string) error {
	broadcasterID, success := GetTwitchUserIDByLogin(channel)
	if !success {
		return nil
	}

	_, err := config.Config.Twitch.Client().CreateEventSubSubscription(&helix.
		EventSubSubscription{
		Type:    eventType,
		Version: "1",
		Condition: helix.EventSubCondition{
			BroadcasterUserID: broadcasterID,
		},
		Transport: helix.EventSubTransport{
			Method:   "webhook",
			Callback: config.Config.Twitch.Webhook.GetURL(channel),
			Secret:   config.Config.Twitch.Webhook.Secret,
		},
	})

	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to send EventSub request")
		return err
	}

	return nil
}

func GetTwitchStreamByLogin(login string) (stream *helix.Stream, err error) {
	streams, err := config.Config.Twitch.Client().GetStreams(&helix.StreamsParams{
		UserLogins: []string{login},
	})

	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to get the stream")
		return nil, err
	}

	if streams.Data.Streams == nil {
		return nil, nil
	}

	return &streams.Data.Streams[0], err
}
