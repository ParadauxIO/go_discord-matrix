package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WebhookMessage struct {
	content    string
	username   string
	avatar_url string
	tts        bool
	embeds     []Embed
}

type Embed struct {
	// TODO
}

type MessageResult struct {
	// TODO
}

func sendMessage(webhook string, message WebhookMessage) MessageResult {
	byteValue, err := json.Marshal(&message)

	if err != nil {
		return MessageResult{} // TODO err obj
	}

	var jsonString string
	json.Unmarshal(byteValue, &jsonString)
	fmt.Printf("%#v", jsonString)

	httpReq, httpErr := http.NewRequest("POST", webhook, bytes.NewBuffer(byteValue))
	httpReq.Header.Set("Content-Type", "application/json")

	if httpErr != nil {
		return MessageResult{} // TODO err obj
	}

	httpClient := &http.Client{}
	httpResp, httpErr := httpClient.Do(httpReq)

	if httpErr != nil {
		return MessageResult{} // TODO err obj
	}

	defer httpResp.Body.Close()

	fmt.Println("response Status:", httpResp.Status)
	fmt.Println("response Headers:", httpResp.Header)
	body, _ := ioutil.ReadAll(httpResp.Body)
	fmt.Println("response Body:", string(body))

	return MessageResult{}

}
