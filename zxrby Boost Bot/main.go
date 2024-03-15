package main

import (
	"BoostTool/Core/Bot"
	"BoostTool/Core/Discord"
	"BoostTool/Core/Utils"
	"time"

	title "github.com/lxi1400/GoTitle"
)


func main() {
	config, _ := Utils.LoadConfig() // Ignoring the error for simplicity

	// Continue with the rest of your program
	Utils.ClearScreen()

	time.Sleep(time.Second * 1)
	Utils.LogInfo("Logging in.", "", "")

	title.SetTitle("Boostility | .gg/boostility | dev_demon1zed & demon1zed.zip")
	Utils.ClearScreen()
	Utils.PrintASCII()

	if config.CustomPersonalization.Onliner {
		Utils.LogInfo("Token Onliner: Enabled", "", "")
		go Discord.Websocket()
	} else {
		Utils.LogInfo("Token Onliner: Disabled", "", "")
	}

	go Bot.Automation()
	Bot.StartBot()
}
