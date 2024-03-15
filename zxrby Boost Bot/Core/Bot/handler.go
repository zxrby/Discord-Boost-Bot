package Bot

import (
	"BoostTool/Core/Discord"
	"BoostTool/Core/Utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)


var EmbedColor, _ = strconv.ParseInt(strings.Replace(config.DiscordSettings.EmbedColor, "#", "", -1), 16, 32)

type Config struct {
	DiscordSettings struct {
		Owners []string `json:"owners"`
	} `json:"discordSettings"`
	PayPalEmail           string `json:"paypal_email"`
	LiveStockMsgChannelID string `json:"live_stock_msg_channelID"`
}

func PreSteps(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var invite string
	var file string
	var err error

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	if !Utils.CheckPermissions(i.Member.User.ID) {
		content := "You do not have permissions to run this command!"
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				{
					Title:       "Bot Error",
					Description: content,
					Color:       int(EmbedColor),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    "Powered by .gg/boostility",
						IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
					},
				},
			},
		})
		return
	}

	servercode := i.Interaction.ApplicationCommandData().Options[0].StringValue()
	amount := i.Interaction.ApplicationCommandData().Options[1].IntValue()
	duration := i.Interaction.ApplicationCommandData().Options[2].IntValue()

	inviteParts := strings.Split(servercode, "/")
	count := len(inviteParts)

	if count == 4 {
		invite = inviteParts[3]
	} else if count == 2 {
		invite = inviteParts[1]
	} else {
		invite = servercode
	}

	if duration == 1 {
		file = "1 Month Tokens.txt"

	} else if duration == 3 {
		file = "3 Month Tokens.txt"
	} else {
		err = errors.New("Choose Proper Duration Type")
	}

	if err != nil {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				{
					Title:       "Bot Error",
					Description: fmt.Sprintf("We have failed to boost a server. The error is below.\n**__%v__**", err.Error()),
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Order Information",
							Value:  fmt.Sprintf("```\nBoost Amount: %v\nServer Invite: %v\n```", amount, invite),
							Inline: false,
						},
					},
					Color: int(EmbedColor),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    "Powered by .gg/boostility",
						IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
					},
				},
			},
		})

		return
	}

	Utils.ClearScreen()
	Utils.PrintASCII()

	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{
			{
				Title:       "Boost Bot",
				Description: "We have started boosting the server provided.",
				Color:       int(EmbedColor),
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "Powered by .gg/boostility",
					IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
				},
			},
		},
	})

	respo, err := Discord.BoostServer(invite, int(amount), file)
	if err != nil {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				{
					Title:       "Bot Error",
					Description: fmt.Sprintf("We have failed to boost a server. The error is below.\n**__%v__**", err.Error()),
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Order Information",
							Value:  fmt.Sprintf("```\nBoost Amount: %v\nServer Invite: %v\n```", amount, invite),
							Inline: false,
						},
					},
					Color: int(EmbedColor),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    "Powered by .gg/boostility",
						IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
					},
				},
			},
		})

		return
	}

	var success string
	if len(respo.SuccessTokens) != 0 {
		success = strings.Join(respo.SuccessTokens, "\n")
	} else {
		success = "No succeeded tokens."
	}

	var failed string
	if len(respo.FailedTokens) != 0 {
		failed = strings.Join(respo.FailedTokens, "\n")
	} else {
		failed = "No failed tokens."
	}
	descrip := fmt.Sprintf("``🔗`` Invite Link: **[%v](https://discord.gg/%v)**\n``💎`` Amount: **%v**\n``📆`` Duration: **%v Month**\n``✅`` Success: **%v**\n``❌`` Failed: **%v**\n``⏰`` Elapsed Time: **%v**", invite, invite, amount, duration, respo.Success, respo.Failed, respo.ElapsedTime)
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{
			{
				Title:       "Boosts Finished",
				Description: descrip,
				Color:       int(EmbedColor),
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "Powered by .gg/boostility",
					IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
				},
			},
		},
	})
	embed := discordgo.MessageEmbed{
		Title:       "Boost Bot",
		Description: fmt.Sprintf("We have boosted a server successfully.\n**Success**: %v | **Failed**: %v\n**Elapsed Time**: %v", respo.Success, respo.Failed, respo.ElapsedTime),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Order Information",
				Value:  fmt.Sprintf("```\nBoost Amount: %v\nServer Invite: %v\n```", amount, invite),
				Inline: false,
			},
			{
				Name:   "``📍`` Succeeded Tokens",
				Value:  fmt.Sprintf("```\n%v\n```", success),
				Inline: false,
			},
			{
				Name:   "``❌`` Failed Tokens",
				Value:  fmt.Sprintf("```\n%v\n```", failed),
				Inline: false,
			},
		},
		Color: int(EmbedColor),
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Powered by .gg/boostility",
			IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
		},
	}
	s.ChannelMessageSendEmbed(config.DiscordSettings.LogsChannel, &embed)

	return
}


// Sample implementation of Cache struct with Get and Set methods
type Cache struct {
	data map[string]string
}

func (c *Cache) Get(key string) (string, error) {
	val, ok := c.data[key]
	if !ok {
		return "", fmt.Errorf("key %s not found", key)
	}
	return val, nil
}

func (c *Cache) Set(key, value string) {
	if c.data == nil {
		c.data = make(map[string]string)
	}
	c.data[key] = value
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func LiveStock(s *discordgo.Session, i *discordgo.InteractionCreate) {
	config, err := LoadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Check if the user invoking the command is one of the owners
	isOwner := false
	for _, ownerID := range config.DiscordSettings.Owners {
		if i.Member.User.ID == ownerID {
			isOwner = true
			break
		}
	}
	if !isOwner {
		// If the user is not an owner, return without executing the command
		fmt.Println("Unauthorized user attempted to use LiveStock command")
		return
	}

	channelID := config.LiveStockMsgChannelID // Obtain the channel ID from the config
	_, err = s.ChannelMessageSendEmbed(i.ChannelID, &discordgo.MessageEmbed{
		Title:       "Successful",
		Description: fmt.Sprintf("Successfully Sent Live Stock to <#%s>", channelID),
		Color:       int(0x00ff00), // Green color
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Powered by .gg/boostility",
			IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
		},
	})

	if err != nil {
		fmt.Println("Error sending stock message:", err)
	}

	sendStockMessage(s, channelID)

	// Create a ticker that ticks every 2 minutes
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	// Run a loop to send the stock message every 2 minutes
	for {
		select {
		case <-ticker.C:
			sendStockMessage(s, channelID)
		}
	}
}

func sendStockMessage(s *discordgo.Session, channelID string) {
	value1 := Utils.Get1mTokens()
	value2 := Utils.Get3MTokens()

	description := fmt.Sprintf("**__1 Month:__**\n`📦` Tokens: **%v**\n`💎` Boosts: **%v**\n\n**__3 Months:__**\n`📦` Tokens: **%v**\n`💎` Boosts: **%v**", value1, value1*2, value2, value2*2)

	// Get the previous message ID
	msg, err := s.ChannelMessages(channelID, 1, "", "", "")
	if err != nil {
		fmt.Println("Error retrieving previous message:", err)
		return
	}

	// If there's a previous message, delete it
	if len(msg) > 0 {
		err = s.ChannelMessageDelete(channelID, msg[0].ID)
		if err != nil {
			fmt.Println("Error deleting previous message:", err)
			return
		}
	}

	// Send the new stock message to the channel
	_, err = s.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Title:       "Live Stock",
		Description: description,
		Color:       int(0x00ff00), // Green color
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Powered by .gg/boostility",
			IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
		},
	})
	if err != nil {
		fmt.Println("Error sending stock message:", err)
	}
}

func stockCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var value1, value2 int

	// Get the value of the "type" option from the interaction
	typet := strings.ToUpper(i.Interaction.ApplicationCommandData().Options[0].StringValue())

	switch typet {
	case "1":
		value1 = Utils.Get1mTokens()
		value2 = value1 * 2 // Boosts for 1 month are double the tokens
	case "3":
		value1 = Utils.Get3MTokens()
		value2 = value1 * 2 // Boosts for 3 months are double the tokens
	case "ALL":
		value1 = Utils.Get1mTokens()
		value2 = Utils.Get3MTokens()
	default:
		// Handle invalid inputs
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Invalid input. Please provide '1', '3', or 'ALL'.",
			},
		})
		return
	}

	var description string
	if typet == "1" || typet == "3" {
		description = fmt.Sprintf("`📦` Tokens: **%v**\n`💎` Boosts: **%v**", value1, value2)
	} else {
		description = fmt.Sprintf("**__1 Month:__**\n`📦` Tokens: **%v**\n`💎` Boosts: **%v**\n\n**__3 Months:__**\n`📦` Tokens: **%v**\n`💎` Boosts: **%v**", value1, value1*2, value2, value2*2)
	}

	// Respond with the stock information
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Boost Bot Stock",
					Description: description,
					Color:       int(EmbedColor),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    "Powered by .gg/boostility",
						IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
					},
				},
			},
		},
	})
}

func restockCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var stat os.FileInfo
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
	typet := i.Interaction.ApplicationCommandData().Options[0].IntValue()
	if !Utils.CheckPermissions(i.Member.User.ID) {
		content := "You do not have permissions to run this command!"
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				{
					Title:       "Bot Error",
					Description: content,
					Color:       int(EmbedColor),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    "Powered by .gg/boostility",
						IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
					},
				},
			},
		})
		return
	}

	var url string
	url = i.Interaction.ApplicationCommandData().Options[1].StringValue()
	resp, err := http.Get("https://paste.ee/d/" + url)
	if err != nil {
		content := "Failed to open file"
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &content,
		})
		return
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	lines := strings.Split(string(b), "\n")

	if typet == 1 {
		stat, _ = os.Stat("./Data/1 Month Tokens.txt")
		for i, s := range lines {
			if i == 0 && stat.Size() != 0 {
				Utils.AppendTextToFile(s+"\n", "1 Month Tokens.txt", "\n")
			} else {
				Utils.AppendTextToFile(s+"\n", "1 Month Tokens.txt")
			}
		}
	} else if typet == 3 {
		stat, _ = os.Stat("./Data/3 Month Tokens.txt")
		for i, s := range lines {
			if i == 0 && stat.Size() != 0 {
				Utils.AppendTextToFile(s+"\n", "3 Month Tokens.txt", "\n")
			} else {
				Utils.AppendTextToFile(s+"\n", "3 Month Tokens.txt")
			}
		}
	}

	content := "Added tokens to the token file successfully"
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &content,
	})
	return
}

func sendCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var file string
	var tokens []string

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	if !Utils.CheckPermissions(i.Member.User.ID) {
		content := "You do not have permissions to run this command!"
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				{
					Title:       "Bot Error",
					Description: content,
					Color:       int(EmbedColor),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    "Powered by .gg/boostility",
						IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
					},
				},
			},
		})
		return
	}
	typed := int(i.Interaction.ApplicationCommandData().Options[0].IntValue())
	amount := int(i.Interaction.ApplicationCommandData().Options[1].IntValue())
	member := i.Interaction.ApplicationCommandData().Options[2].UserValue(s).ID

	_, err := os.Create("./Data/tokens.txt")
	if err != nil {
		Utils.LogError(err.Error(), "", "")
		return
	}

	file123, err := os.Open("./Data/tokens.txt")
	if err != nil {
		Utils.LogError(err.Error(), "", "")
		return
	}

	if typed == 3 {
		file = "3 Month Tokens.txt"
	} else if typed == 1 {
		file = "1 Month Tokens.txt"
	}

	for i := 0; i < amount; i++ {
		token12 := Utils.SendToken(file)
		tokens = append(tokens, token12+"\n")
	}

	for _, tokens1 := range tokens {
		Utils.AppendTextToFile(tokens1, "tokens.txt")
	}

	fileToSend := &discordgo.File{
		Name:        "tokens.txt", // Name of the file when received by the user
		Reader:      file123,
		ContentType: "text/plain", // Adjust the content type as needed
	}

	channel, err := s.UserChannelCreate(member)
	channelid := channel.ID

	_, err = s.ChannelMessageSendComplex(channelid, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Bot Tokens",
			Description: fmt.Sprintf("`📦` Tokens: **%v**", amount),
			Color:       int(EmbedColor),
			Footer: &discordgo.MessageEmbedFooter{
				Text:    "Powered by .gg/boostility",
				IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
			},
		},
		Files: []*discordgo.File{fileToSend},
	})

	if err != nil {
		Utils.LogError(err.Error(), "", "")
		content := "Failed Sending Tokens, User has DM's Disabled. Tokens Returned to File!"
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &content,
		})

		for _, tokens1 := range tokens {
			Utils.AppendTextToFile(tokens1, file)
		}

	} else {
		content := "Successfully Sent Tokens!"
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &content,
		})

	}

	_ = file123.Close()

	err = os.Remove("./Data/tokens.txt")
	if err != nil {
		Utils.LogError(err.Error(), "", "")
	}

	return
}

func helpCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Create an embed message with information about available commands
	embed := &discordgo.MessageEmbed{
		Title:       "Bot Commands",
		Description: "Here are the available commands and their descriptions:",
		Color:       0x00ff00, // Green color
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Powered by .gg/boostility",
			IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "boost",
				Value:  "Boost a Server\nUsage: `/boost <invite> <amount> <months>`\nExample: `/boost ABC123 2 3`",
				Inline: false,
			},
			{
				Name:   "stock",
				Value:  "Boost Bot Stock\nUsage: `/stock <type>`\nExample: `/stock 1`",
				Inline: false,
			},
			{
				Name:   "restock",
				Value:  "Add tokens to the tokens file\nUsage: `/restock <type> <code>`\nExample: `/restock 1 https://paste.ee/p/xxxxx`",
				Inline: false,
			},
			{
				Name:   "send",
				Value:  "Send tokens from the token list\nUsage: `/send <type> <amount> <recipient>`\nExample: `/send 1 5 @RecipientUser`",
				Inline: false,
			},
			{
				Name:   "paypal",
				Value:  "Sends My PayPal So You Can Send Me Money\nUsage: `/paypal`",
				Inline: false,
			},
		},
	}

	// Send the embed message to the user who triggered the command
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
	if err != nil {
		Utils.LogError("Error sending help message", "", err.Error())
	}
}

func restart(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Check if the user invoking the command is one of the owners
	isOwner := false
	for _, ownerID := range config.DiscordSettings.Owners {
		if i.Member.User.ID == ownerID {
			isOwner = true
			break
		}
	}
	if !isOwner {
		// If the user is not an owner, return without executing the command
		fmt.Println("Unauthorized user attempted to use restart command")
		return
	}
	// Send message indicating restart
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Restarting...",
		},
	})

	// Start the new process
	cmd := exec.Command("cmd", "/C", "boostility.exe")
	err := cmd.Start()
	if err != nil {
		// Handle error
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Error: " + err.Error(),
			},
		})
		return
	}

	// Terminate the current process
	os.Exit(0)
}

func shutdown(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Check if the user invoking the command is one of the owners
	isOwner := false
	for _, ownerID := range config.DiscordSettings.Owners {
		if i.Member.User.ID == ownerID {
			isOwner = true
			break
		}
	}
	if !isOwner {
		// If the user is not an owner, return without executing the command
		fmt.Println("Unauthorized user attempted to use Shutdown command")
		return
	}
	// Send message indicating shutdown
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Shutting down...",
		},
	})

	// Cleanly close down the Discord session
	s.Close()

	// Terminate the application
	os.Exit(0)
}

func PayPalCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Read the config file
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// Parse the config data into a Config struct
	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return
	}

	// Extract the amount from the command options
	amount := i.Interaction.ApplicationCommandData().Options[0].StringValue()

	// Create the PayPal embed message
	embed := &discordgo.MessageEmbed{
		Title: "PayPal Information",
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Powered by .gg/boostility",
			IconURL: "https://images-ext-2.discordapp.net/external/77SrO1XOb4psFtHKvW9i-07cXuMEagxwAfYFUQu_gz0/%3Fsize%3D1024/https/cdn.discordapp.com/icons/1203345752093491250/a_fc3fae24d895d84c0bf995a55e301063.gif",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "`Email`",
				Value:  fmt.Sprintf("**%s**\n", config.PayPalEmail),
				Inline: false,
			},
			{
				Name:   "`Method`",
				Value:  "**Friends & Family**\n",
				Inline: false,
			},
			{
				Name:   "`Amount`",
				Value:  fmt.Sprintf("**$%s**\n", amount),
				Inline: false,
			},
			{
				Name:   "`Disclaimer`",
				Value:  "**Send As Goods And Services Will Equal To You Not Getting A Refund + No Items And We Do Not Cover Fees/Taxes**\n",
				Inline: false,
			},
		},
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
	if err != nil {
		fmt.Println("Error responding to /paypal command:", err)
	}
}
