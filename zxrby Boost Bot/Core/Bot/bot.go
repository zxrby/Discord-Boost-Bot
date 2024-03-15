package Bot

import (
	"BoostTool/Core/Utils"
	"flag"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session
var config, _ = Utils.LoadConfig()

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "boost",
			Description: "Boost a Server",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "invite",
					Description: "Invite Code of Server",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "amount",
					Description: "Amount of Boosts (Even Number Only)",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "months",
					Description: "Number of Months (1 or 3)",
					Required:    true,
				},
			},
		},
		{
			Name:        "stock",
			Description: "Boost Bot Stock",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "type",
					Description: "Type of Tokens",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "1",
							Value: "1",
						},
						{
							Name:  "3",
							Value: "3",
						},
						{
							Name:  "ALL",
							Value: "ALL",
						},
					},
				},
			},
		},
		{
			Name:        "restock",
			Description: "Add tokens to the tokens file",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "type",
					Description: "Type of Tokens (1 or 3)",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "code",
					Description: "Use Code From Paste.ee URl (Ex. https://paste.ee/p/xxxxx)",
					Required:    true,
				},
			},
		},
		{
			Name:        "send",
			Description: "Send tokens from the token list",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "type",
					Description: "Type of Tokens (1 or 3)",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "amount",
					Description: "The amount of tokens",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "recipient",
					Description: "Who are you Sending the Tokesn to?",
					Required:    true,
				},
			},
		},
		{
			Name:        "paypal",
			Description: "Send PayPal Information",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "amount",
					Description: "The amount of money",
					Required:    true,
				},
			},
		},
		{
			Name:        "help",
			Description: "Sends You Help!",
		},
		{
			Name:        "live-stock",
			Description: "shows live stock",
		},
		{
			Name:        "restart",
			Description: "restarts bot",
		},
		{
			Name:        "shutdown",
			Description: "shutdsdown bot",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"boost":      PreSteps,
		"stock":      stockCommand,
		"restock":    restockCommand,
		"send":       sendCommand,
		"help":       helpCommand,
		"paypal":     PayPalCommand,
		"live-stock": LiveStock,
		"restart":    restart,
		"shutdown":   shutdown,
	}
)

func init() {
	var err error
	s, err = discordgo.New("Bot " + config.DiscordSettings.Token)
	if err != nil {
		Utils.LogError("Invalid Bot Parameters", "Error", err.Error())
		return
	}

	if config.DiscordSettings.BotStatus != "" {
		s.Identify.Presence.Game = discordgo.Activity{
			Name: config.DiscordSettings.BotActivity,
			Type: getActivityType(),
		}
	}

}

func getActivityType() discordgo.ActivityType {
	switch config.DiscordSettings.BotStatus {
	case "playing":
		return discordgo.ActivityTypeGame
	case "watching":
		return discordgo.ActivityTypeWatching
	case "listening":
		return discordgo.ActivityTypeListening
	case "competing":
		return discordgo.ActivityTypeCompeting
	default:
		return discordgo.ActivityTypeWatching
	}
}

func StartBot() {
	var err error
	RemoveCommands := flag.Bool("rmcmd", true, "Remove all commands after shutting down or not")

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		botd := s.State.User.Username + "#" + s.State.User.Discriminator
		Utils.LogInfo("Successfully Logged In", "Bot", botd)
	})
	err = s.Open()
	if err != nil {
		Utils.LogError(err.Error(), "", "")
		return
	}
	Utils.LogInfo("Registering Commands to Guild", "Guild", config.DiscordSettings.GuildID)
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err1 := s.ApplicationCommandCreate(s.State.User.ID, config.DiscordSettings.GuildID, v)
		if err1 != nil {
			Utils.LogError(err1.Error(), "", "")
		} else {
			registeredCommands[i] = cmd
		}
	}
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	Utils.LogInfo("Press Ctrl+C to Shutdown Commands", "", "")
	<-stop

	if *RemoveCommands {
		Utils.LogInfo("Shutting Down Commands...", "", "")

		for _, v := range registeredCommands {
			if v != nil {
				_ = s.ApplicationCommandDelete(s.State.User.ID, config.DiscordSettings.GuildID, v.ID)
			}
		}
	}
}
