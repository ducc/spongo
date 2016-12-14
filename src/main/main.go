package main

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"strings"
	"regexp"
	"strconv"
	"errors"
	"math/rand"
	"time"
	"fmt"
)

const (
	no_permission_msg = "You do not have permission to run this command! Required permission level: `%d`"
	config_path        = "config.toml"
	started_msg_format = "spongo running! usr: %s#%s (%s)\n"
)

type (
	config struct {
		Token  string
		Shards int
		Owner  string
		Prefix string
	}

	context struct {
		session *discordgo.Session
		event   *discordgo.MessageCreate
		guild   *discordgo.Guild
		channel *discordgo.Channel
		content string
		args    []string
	}

	command func(*context)
)

func (ctx *context) reply(text string) *discordgo.Message {
	msg, err := ctx.session.ChannelMessageSend(ctx.event.ChannelID, text)
	if err != nil {
		log.Println(err)
		return nil
	}
	return msg
}

func (ctx *context) invalidArgs(usage string) *discordgo.Message {
	return ctx.reply("Invalid arguments! Usage: " + usage)
}

func (ctx *context) noPermission(level int) *discordgo.Message {
	return ctx.reply(fmt.Sprintf(no_permission_msg, level))
}

func (ctx *context) err(text string, err error) *discordgo.Message {
	log.Println("Error during command processing,", text, err)
	return ctx.reply("An error occured! " + text)
}

var (
	conf     *config
	localId  string
	commands = make(map[string]command)
	parseMentionRegex *regexp.Regexp
	parseUserRegex *regexp.Regexp
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
	registerCommands()
	commands["help"] = func(ctx *context) {
		buff := bytes.Buffer{}
		for name := range commands {
			buff.WriteString(name)
			buff.WriteString(", ")
		}
		ctx.reply(strings.TrimSuffix(buff.String(), ", "))
	}
	parseUserRegex, _ = regexp.Compile("(.{2,32})#(\\d{4})")
	parseMentionRegex, _ = regexp.Compile("<@\\!?(\\d+)>")
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
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println("Could not get source channel,", err)
		return
	}
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		log.Println("Could not get source guild,", err)
		return
	}
	ctx := &context{
		session: s,
		event:   m,
		guild: guild,
		channel: channel,
		content: content,
		args:    args[1:],
	}
	cmd(ctx)
}

func loadConfig(file string) (*config, error) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var conf config
	_, err = toml.Decode(string(body), &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

func parseMember(guild *discordgo.Guild, input string) (*discordgo.Member, error) {
	var t uint8
	var id string
	var username string
	var discrim string
	if _, err := strconv.Atoi(input); err == nil {
		t = 0
		id = input
	} else if !strings.Contains(input, "#") {
		if strings.Contains(input, "<") {
			t = 2
			matches := parseMentionRegex.FindAllStringSubmatch(input, -1)
			id = matches[0][1]
		} else {
			t = 1
			username = input
		}
	} else {
		t = 3
		matches := parseUserRegex.FindAllStringSubmatch(input, -1)
		username = matches[0][1]
		discrim = matches[0][2]
	}
	for _, m := range guild.Members {
		if t == 0 || t == 2 {
			if m.User.ID == id {
				return m, nil
			}
		} else if (t == 3) {
			if m.User.Username == username && m.User.Discriminator == discrim {
				return m, nil
			}
		} else {
			if m.User.Username == username {
				return m, nil
			}
		}
	}
	return nil, errors.New("Could not find member")
}

func randomBool() bool {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 2 == 0
}

func randomIntInRange(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func registerCommands() {
	info()
	util()
	mod()
	social()
	misc()
	game()
	admin()
}
