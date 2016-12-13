package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

type (
	context struct {
		session *discordgo.Session
		event   *discordgo.MessageCreate
		content string
		args    []string
	}

	command func(*context)
)

const (
	config_path        = "config.toml"
	started_msg_format = "spongo running! usr: %s#%s (%s)\n"
)

var (
	conf    *config
	localId string
	commands = make(map[string]command)
)

func main() {
	var err error
	conf, err = loadConfig(config_path)
	if err != nil {
		log.Fatal(err)
		return
	}
	discord, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		log.Fatal(err)
		return
	}
	localUser, err := discord.User("@me")
	if err != nil {
		log.Fatal(err)
		return
	}
	localId = localUser.ID
	discord.AddHandler(messageCreate)

	commands["dog"] = func(ctx *context) {
		ctx.session.ChannelMessageSend(ctx.event.ChannelID, "Big dog :)")
	}

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf(started_msg_format, localUser.Username, localUser.Discriminator, localId)
	<-make(chan struct{})
	return
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == localId {
		return
	}
	if len(m.Content) <= len(conf.Prefix) {
		return
	}
	if !strings.HasPrefix(m.Content, conf.Prefix) {
		return
	}
	content := strings.TrimPrefix(m.Content, conf.Prefix)
	if content[0] == ' ' {
		content = content[1:]
	}
	args := strings.Split(content, " ")
	cmd, ok := commands[args[0]]
	if !ok {
		return
	}
	ctx := &context{
		session: s,
		event: m,
		content: content,
		args: args[1:],
	}
	cmd(ctx)
}
