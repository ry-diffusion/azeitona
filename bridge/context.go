package bridge

type Context struct {
	DiscordToken        string
	IRCServer           string
	DiscordChannel      string
	DiscordMessages     chan Message
	DiscordWebhookId    string
	DiscordWebhookToken string
	IRCMessages         chan Message
	IRCNickname         string
	IRCChannel          string
}
