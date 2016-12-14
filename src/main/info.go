package main

func info() {
	commands["ping"] = func(ctx *context) {
		ctx.reply("Pong!")
	}
	commands["invite"] = func(ctx *context) {
		ctx.reply("BOYYY THIS IS A FUCKIN INVITE LINK!!!")
	}
	commands["stats"] = func(ctx *context) {
		ctx.reply("here are some stats lol")
	}
}
