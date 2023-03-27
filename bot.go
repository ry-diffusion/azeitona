package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/milyth/azeitona/bridge"
)

func main() {
	discordToken := os.Getenv("DISCORD_TOKEN")
	ircAddress := os.Getenv("IRC_SERVER")

	fmt.Printf("IRC: %s\n", ircAddress)
	ctx := bridge.Context{
		DiscordToken:        discordToken,
		IRCServer:           ircAddress,
		IRCNickname:         os.Getenv("IRC_NICKNAME"),
		DiscordChannel:      os.Getenv("DISCORD_CHANNEL"),
		IRCChannel:          os.Getenv("IRC_CHANNEL"),
		DiscordWebhookId:    os.Getenv("DISCORD_WEBHOOK_ID"),
		DiscordWebhookToken: os.Getenv("DISCORD_WEBHOOK_TOKEN"),
		DiscordMessages:     make(chan bridge.Message),
		IRCMessages:         make(chan bridge.Message),
	}

	if err := bridge.Discord(ctx); err != nil {
		fmt.Printf("failed to initialize discord bridge: %v\n", err)
	}

	if err := bridge.IRC(ctx); err != nil {
		fmt.Printf("unable to initialize IRC bridge: %v\n", err)
	}

	for {
	}
}
