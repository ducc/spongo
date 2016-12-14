package main

import "os"

func admin() {
	commands["admin"] = func(ctx *context) {
		if ctx.event.Author.ID != conf.Owner {
			ctx.noPermission(2)
			return
		}
		if len(ctx.args) == 0 {
			ctx.invalidArgs("`admin <subcommand>`.")
			return
		}
		switch ctx.args[0] {
		case "stop":
			ctx.reply("Bye :sob:")
			os.Exit(0)
			break
		}
	}
}
