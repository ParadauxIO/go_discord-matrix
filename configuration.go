package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var config Configuration

type Configuration struct {
	Mappings             []Mapping            `json:"mappings"`
	DiscordConfiguration DiscordConfiguration `json:"discord"`
	MatrixConfiguration  MatrixConfiguration  `json:"matrix"`
}

type DiscordConfiguration struct {
	Token string `json:"token"`
}

type MatrixConfiguration struct {
	HomeServer string `json:"home-server"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Mappings struct {
	Mappings []Mapping `json:"mappings"`
}

type Mapping struct {
	DiscordChannelWebhook string `json:"discord-webhook"`
}

func loadConfig() {
	cFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println("There was an error loading the configuration.")
		fmt.Println(err)
		os.Exit(1)
	}

	defer cFile.Close()

	config = parseConfig(cFile)
	printConfig()
}

func parseConfig(configFile *os.File) Configuration {
	byteValue, err := ioutil.ReadAll(configFile)

	if err != nil {
		fmt.Println("There was an error parsing the configuration.")
		fmt.Println(err)
		os.Exit(1)
	}

	var c Configuration

	json.Unmarshal(byteValue, &c)

	return c
}

func printConfig() {
	fmt.Printf("%#v", config)
}
