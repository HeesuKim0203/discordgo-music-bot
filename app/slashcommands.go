package app

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	RemoveCommands                 = true
	GuildID                        = ""
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer
	//integerOptionMinValue          = 1.0

	commands = []*discordgo.ApplicationCommand{
		{
			Name:                     "inviataion-channel",
			Description:              "Invitation channel command",
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermission,
		},
		// {
		// 	Name:                     "play-song",
		// 	Description:              "play song",
		// 	DefaultMemberPermissions: &defaultMemberPermissions,
		// 	DMPermission:             &dmPermission,
		// },
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"inviataion-channel": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			_, err := s.ChannelVoiceJoin(i.GuildID, i.ChannelID, false, true)

			if err != nil {
				log.Println("Falied Channel Invitation!")
				return
			} else {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Invitation Successful!",
					},
				})
			}
		},
		// "search-song": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		// },
	}
)

func RegisterCommends(discord *discordgo.Session) []*discordgo.ApplicationCommand {
	// Add slash command
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {

		cmd, err := discord.ApplicationCommandCreate(discord.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	return registeredCommands
}

func UnRegister(discord *discordgo.Session, registeredCommands []*discordgo.ApplicationCommand) {
	if RemoveCommands {
		log.Println("Removing commands...")
		for _, v := range registeredCommands {
			err := discord.ApplicationCommandDelete(discord.State.User.ID, GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}
	log.Println("Gracefully shutting down.")
}
