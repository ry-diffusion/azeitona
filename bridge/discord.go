package bridge

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Discord(ctx Context) error {
	discord, err := discordgo.New("Bot " + ctx.DiscordToken)

	if err != nil {
		return fmt.Errorf("unable to create discord context: %v", err)
	}

	discord.AddHandler(func(s *discordgo.Session, msg *discordgo.MessageCreate) {
		if msg.WebhookID == ctx.DiscordWebhookId || msg.ChannelID != ctx.DiscordChannel {
			return
		}

		for _, line := range strings.Split(msg.Content, "\n") {
			if len(line) > 0 {
				ctx.DiscordMessages <- Message{
					Author:  msg.Author.String(),
					Content: line,
				}
			}
		}

		for _, attach := range msg.Attachments {
			ctx.DiscordMessages <- Message{
				Author:  msg.Author.String(),
				Content: attach.URL,
			}
		}
	})

	discord.AddHandler(func(s *discordgo.Session, _ *discordgo.Ready) {
		fmt.Println("Discord bridge is ready")
		go func() {
			for {
				ircMessage := <-ctx.IRCMessages
				content := strings.ReplaceAll(ircMessage.Content, "@", "")
				if len(content) < 1 {
					continue
				}

				_, err := s.WebhookExecute(ctx.DiscordWebhookId, ctx.DiscordWebhookToken, true, &discordgo.WebhookParams{
					Content:  content,
					Username: ircMessage.Author,
				})

				if err != nil {
					fmt.Printf("unable to send message: %v", err)
				}
			}
		}()
	})

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	return discord.Open()
}
