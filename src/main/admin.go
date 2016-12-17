package main

import (
	"os"
	"github.com/robertkrimen/otto"
	"strings"
)

var jsVm = otto.New()

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
		case "eval": {
			input := strings.Join(ctx.args[1:], " ")
			jsVm.Set("ctx", ctx)
			val, err := jsVm.Run(input)
			if err != nil {
				ctx.reply(err.Error())
				break
			}
			if !val.IsNull() {
				str, err := val.ToString()
				if err != nil {
					ctx.reply(err.Error())
					break
				}
				ctx.reply(str)
			}
			break
		}
		}
	}
}
