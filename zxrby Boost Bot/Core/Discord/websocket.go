package Discord

import (
	"BoostTool/Core/Utils"
	"math/rand"
	"strings"
	"time"

	//"time"

	"github.com/gorilla/websocket"
)

var randomStatus = []string{"online", "idle", "dnd"}
var types = []string{"Playing", "Streaming", "Watching", "Listening"}
var games = []string{"Minecraft", "Fortnite", "Roblox", "The Elder Scrolls: Online", "DCS World Steam Edit"}
var music = []string{"Spotify", "Deezer", "Apple Music", "YouTube", "SoundCloud", "Pandora", "Tidal", "Amazon Music", "Google Play Music", "Apple Podcasts", "iTunes", "Beatport"}
var watch = []string{"YouTube", "Twitch", "Kick"}
var game string
var gamejson map[string]interface{}

func Websocket() {
	for {
		tokens := Utils.OnlinerTokens()
		for _, line := range tokens {
			if line != "" {
				go func(line string) {
					KeepOnline(line)
				}(line)
			}
		}
		time.Sleep(time.Millisecond * 41250)
	}
}

func KeepOnline(token string) {
	var auth map[string]interface{}
	mutex.Lock()
	if strings.Contains(token, ":") {
		token = strings.Split(token, ":")[2]
	}

	stype := types[rand.Intn(len(types))]

	if stype == "Playing" {
		game = games[rand.Intn(len(games))]
		gamejson = map[string]interface{}{
			"name": game,
			"type": 0,
		}
	} else if stype == "Streaming" {
		game = games[rand.Intn(len(games))]
		gamejson = map[string]interface{}{
			"name": game,
			"type": 1,
			"url":  "https://www.twitch.tv/kaicenat",
		}
	} else if stype == "Listening" {
		game = music[rand.Intn(len(music))]
		gamejson = map[string]interface{}{
			"name": game,
			"type": 2,
		}
	} else if stype == "Watching" {
		game = watch[rand.Intn(len(watch))]
		gamejson = map[string]interface{}{
			"name": game,
			"type": 3,
		}
	}
	ws, _, err := websocket.DefaultDialer.Dial("wss://gateway.discord.gg/encoding=json&v=9&compress=zlib-stream", nil)

	if err != nil {
		Utils.LogError(err.Error(), "", "")
	}

	if len(config.CustomPersonalization.Status) != 0 || len(config.CustomPersonalization.StatusEmoji) != 0 {
		status := config.CustomPersonalization.Status[rand.Intn(len(config.CustomPersonalization.Status))]
		emoji := config.CustomPersonalization.StatusEmoji[rand.Intn(len(config.CustomPersonalization.StatusEmoji))]
		auth = map[string]interface{}{
			"op": 2,
			"d": map[string]interface{}{
				"token":        token,
				"capabilities": 61,
				"properties": map[string]interface{}{
					"os":                  "Windows",
					"browser":             "Chrome",
					"device":              "",
					"system_locale":       "en-US",
					"browser_user_agent":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
					"browser_version":     "118.0.0.0",
					"os_version":          "10",
					"release_channel":     "stable",
					"client_build_number": 226220,
				},
				"presence": map[string]interface{}{
					"status": randomStatus[rand.Intn(len(randomStatus))],
					"game": map[string]interface{}{
						"name":  "null",
						"state": status,
						"emoji": map[string]interface{}{
							"id":       nil,
							"name":     emoji,
							"animated": false,
						},
						"type": 4,
					},
					"afk": false,
				},
				"compress": false,
				"client_state": map[string]interface{}{
					"guild_hashes":                map[string]interface{}{},
					"highest_last_message_id":     "0",
					"read_state_version":          0,
					"user_guild_settings_version": -1,
				},
			},
		}
	} else {
		auth = map[string]interface{}{
			"op": 2,
			"d": map[string]interface{}{
				"token": token,
				"properties": map[string]interface{}{
					"os":                  "Windows",
					"browser":             "Chrome",
					"device":              "",
					"system_locale":       "en-US",
					"browser_user_agent":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
					"browser_version":     "118.0.0.0",
					"os_version":          "10",
					"release_channel":     "stable",
					"client_build_number": 226220,
				},
				"presence": map[string]interface{}{
					"game":   gamejson,
					"status": randomStatus[rand.Intn(len(randomStatus))],
					"since":  0,
					"afk":    false,
				},
				"compress": false,
				"client_state": map[string]interface{}{
					"guild_hashes":                map[string]interface{}{},
					"highest_last_message_id":     "0",
					"read_state_version":          0,
					"user_guild_settings_version": -1,
				},
			},
		}
	}

	err = ws.WriteJSON(auth)
	if err != nil {
		Utils.LogError(err.Error(), "", "")
	}

	heartbeatJSON := map[string]interface{}{
		"op":    1,
		"token": "",
		"d":     nil,
	}

	err = ws.WriteJSON(heartbeatJSON)

	if err != nil {
		Utils.LogError(err.Error(), "", "")
	}
	mutex.Unlock()
}
