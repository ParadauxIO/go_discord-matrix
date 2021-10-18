package main

import "fmt"

func main() {
	fmt.Println("Matrix <-> Discord Bridge")

	fmt.Println("... loading configuration")
	//loadConfig()

	sendMessage(webhook, WebhookMessage{
		Content:   "This is a test",
		Username:  "Bridge",
		AvatarUrl: "https://pbs.twimg.com/profile_images/510528591424024576/Fzk0vaGr.png",
		Tts:       false,
		Embeds:    []Embed{},
	})
}
