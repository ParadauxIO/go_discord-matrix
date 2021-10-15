package main

type WebhookMessage struct {
	Content   string
	Username  string
	AvatarUrl string
	TTS       bool
	Embeds    []Embed
}

type Embed struct {
	// TODO
}

type MessageResult struct {
	// TODO
}

func sendMessage(webhook string, message WebhookMessage) MessageResult {

	return MessageResult{}

}
