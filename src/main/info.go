package main

import (
	"runtime"
	"github.com/bwmarrin/discordgo"
	"time"
	"fmt"
	"github.com/dustin/go-humanize"
)

const stats_format = "```go version: %s\ndiscordgo version: %s\nuptime: %s\nmemory used: %s / %s (%s garbage collected)\nconcurrent tasks: %d\ncurrent shard: %d\nshard count: %d```"

var startTime = time.Now()

func info() {
	commands["ping"] = func(ctx *context) {
		ctx.reply("Pong!")
	}
	commands["invite"] = func(ctx *context) {
		ctx.reply("Invite the bot using this link: " + conf.Invite)
	}
	commands["stats"] = func(ctx *context) {
		uptime := formatDurationString(time.Now().Sub(startTime))
		stats := runtime.MemStats{}
		runtime.ReadMemStats(&stats)
		ctx.reply(fmt.Sprintf(stats_format, runtime.Version(), discordgo.VERSION, uptime, humanize.Bytes(stats.Alloc),
			humanize.Bytes(stats.Sys), humanize.Bytes(stats.TotalAlloc), runtime.NumGoroutine(), ctx.session.ShardID,
			ctx.session.ShardCount))
	}
}
