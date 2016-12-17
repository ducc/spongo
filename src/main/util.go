package main

import "fmt"

const (
	userinfo_format = "User: %s#%s\nNickname: %s\nID: %s\nAvatar: https://discordapp.com/api/users/%s/avatars/%s.jpg"
	channelinfo_format = "Name: %s\nID: %s\nChannel topic: %s"
	serverinfo_format = "Name: %s\nID: %s\nRegion: %s\nMembers: %d\nChannels: %d\nOwner: %s#%s\nIcon: https://discordapp.com/api/guilds/%s/icons/%s.jpg"
)

func util() {
	commands["userinfo"] = func(ctx *context) {
		if len(ctx.args) == 0 {
			ctx.invalidArgs("`userinfo <user>`\nuser can be `username#discrim` format, `user id` or just `username`.")
			return
		}
		member, err := parseMember(ctx.guild, ctx.args[0]);
		if err != nil {
			ctx.reply("Invalid user `" + ctx.args[0] + "`.")
			return
		}
		ctx.reply(fmt.Sprintf(userinfo_format, member.User.Username, member.User.Discriminator, member.Nick,
			member.User.ID, member.User.ID, member.User.Avatar))
	}
	commands["channelinfo"] = func(ctx *context) {
		ctx.reply(fmt.Sprintf(channelinfo_format, ctx.channel.Name, ctx.channel.ID, ctx.channel.Topic))
	}
	commands["serverinfo"] = func(ctx *context) {
		user, err := ctx.session.User(ctx.guild.OwnerID)
		if err != nil {
			ctx.reply("Could not find guild owner!")
			return
		}
		ctx.reply(fmt.Sprintf(serverinfo_format, ctx.guild.Name, ctx.guild.ID, ctx.guild.Region, len(ctx.guild.Members),
			len(ctx.guild.Channels), user.Username, user.Discriminator, ctx.guild.ID, ctx.guild.Icon))
	}
	commands["serverlist"] = func(ctx *context) {

	}
	commands["directory"] = func(ctx *context) {

	}
	commands["convert"] = func(ctx *context) {

	}
}
