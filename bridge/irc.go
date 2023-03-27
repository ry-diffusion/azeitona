package bridge

import (
	"fmt"

	irc "github.com/thoj/go-ircevent"
)

func IRC(ctx Context) error {
	conn := irc.IRC(ctx.IRCNickname, ctx.IRCNickname)
	conn.AddCallback("PRIVMSG", func(e *irc.Event) {
		ctx.IRCMessages <- Message{
			Author:  e.Nick,
			Content: e.Message(),
		}
	})

	conn.AddCallback("001", func(e *irc.Event) {
		conn.Join(ctx.IRCChannel)
		go func() {
			for {
				msg := <-ctx.DiscordMessages
				conn.Privmsgf(ctx.IRCChannel, "[%s] %s", msg.Author, msg.Content)
			}
		}()
	})

	err := conn.Connect(ctx.IRCServer)

	if err != nil {
		return fmt.Errorf("unable to connect: %v", err)
	}

	go conn.Loop()

	return nil
}
