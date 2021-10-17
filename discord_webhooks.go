package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WebhookMessage struct {
	Content   string  `json:"content"`
	Username  string  `json:"username"`
	AvatarUrl string  `json:"avatar_url"`
	Tts       bool    `json:"tts"`
	Embeds    []Embed `json:"embeds"`
}

type Embed struct {
	// TODO
}

type MessageResult struct {
	// TODO
}

func sendMessage(webhook string, message WebhookMessage) {
	byteValue, err := json.Marshal(&message)

	if err != nil {
		//	return MessageResult{} // TODO err obj
		return
	}

	var jsonString string
	json.Unmarshal(byteValue, &jsonString)
	fmt.Printf("%#v", jsonString)

	httpReq, httpErr := http.NewRequest("POST", webhook, bytes.NewBuffer(byteValue))
	httpReq.Header.Set("Content-Type", "application/json")

	if httpErr != nil {
		//return MessageResult{} // TODO err obj
		return
	}

	httpClient := &http.Client{}
	httpResp, httpErr := httpClient.Do(httpReq)

	if httpErr != nil {
		//	return MessageResult{} // TODO err obj
		return
	}

	defer httpResp.Body.Close()

	// fmt.Println("response Status:", httpResp.Status)
	// fmt.Println("response Headers:", httpResp.Header)
	// body, _ := ioutil.ReadAll(httpResp.Body)
	// fmt.Println("response Body:", string(body))

	// return MessageResult{}

}
