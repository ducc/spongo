package main

import (
	"fmt"
	"strconv"
)

const (
	roll_format = "Rolled a `%d`!"
)

var (
	eightball_responses = [20]string{
		"It is certain", "It is decidedly so", "Without a doubt", "Yes, definitely", "You may rely on it",
		"As I see it, yes", "Most likely", "Outlook good", "Yes", "Signs point to yes", "Reply hazy try again",
		"Ask again later", "Better not tell you", "Cannot predict now", "Concentrate and ask again",
		"Don't count on it", "My reply is no", "My sources say no", "Outlook not so good", "Very doubtful",
	}
)

func misc() {
	commands["coinflip"] = func(ctx *context) {
		if randomBool() {
			ctx.reply("Heads")
		} else {
			ctx.reply("Tails")
		}
	}
	commands["roll"] = func(ctx *context) {
		var max int
		if len(ctx.args) == 1 {
			var err error
			max, err = strconv.Atoi(ctx.args[0])
			if err != nil {
				ctx.invalidArgs("`roll [max number]`\nDefault max number: `100`")
				return
			}
		} else {
			max = 100
		}
		result := randomIntInRange(0, max)
		ctx.reply(fmt.Sprintf(roll_format, result))
	}
	commands["choice"] = func(ctx *context) {

	}
	commands["8ball"] = func(ctx *context) {
		num := randomIntInRange(0, 19)
		ctx.reply(eightball_responses[num])
	}
	commands["google"] = func(ctx *context) {

	}
	commands["cat"] = func(ctx *context) {

	}
	commands["startpoll"] = func(ctx *context) {

	}
	commands["viewpoll"] = func(ctx *context) {

	}
	commands["endpoll"] = func(ctx *context) {

	}
	commands["vote"] = func(ctx *context) {

	}
	commands["twitchemote"] = func(ctx *context) {

	}
}
