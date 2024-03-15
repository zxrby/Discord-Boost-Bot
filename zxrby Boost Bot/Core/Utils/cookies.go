package Utils

import (
	"encoding/base64"
	"fmt"
)

func SuperProperties() string {
	//properties := map[string]interface{}{
	//	"os":                  "Windows",
	//	"browser":             "Discord Client",
	//	"release_channel":     "stable",
	//	"client_version":      "1.0.9020",
	//	"os_version":          "10.0.19044",
	//	"os_arch":             "x64",
	//	"app_arch":            "ia32",
	//	"system_locale":       "en-US",
	//	"browser_user_agent":  "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9020 Chrome/108.0.5359.215 Electron/22.3.26 Safari/537.36",
	//	"browser_version":     "22.3.26",
	//	"client_build_number": 239468,
	//	"native_build_number": 38517,
	//	"client_event_source": nil,
	//	"design_id":           0,
	//}

	//propertiesJSON, err := json.Marshal(properties)
	//if err != nil {
	//	fmt.Println("Error encoding properties to JSON:", err)
	//
	//}
	properties := fmt.Sprintf(`{"os":"Windows","browser":"Chrome","device":"","system_locale":"en-GB","browser_user_agent":"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.1023 Chrome/98.0.4758.141 Electron/17.4.9 Safari/537.36","browser_version":"109.0.0.0","os_version":"10","referrer":"","referring_domain":"","referrer_current":"","referring_domain_current":"","release_channel":"stable","client_build_number":171027,"client_event_source":null}`)
	propertiesEncoded := base64.StdEncoding.EncodeToString([]byte(properties))
	return propertiesEncoded
}

//func GetCookies(client tls_client.HttpClient) (string, string, string) {
//	req, err := http.NewRequest(http.MethodGet, "https://discord.com", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	headers := http.Header{
//		"accept":          {"*/*"},
//		"accept-encoding": {"gzip"},
//		"accept-language": {"en-US,en-NZ;q=0.9"},
//		"content-type":    {"application/json"},
//		"origin":          {"https://discord.com"},
//		"referer":         {"https://google.com"},
//		"sec-fetch-dest":  {"empty"},
//		"sec-fetch-mode":  {"cors"},
//		"sec-fetch-site":  {"same-origin"},
//		"user-agent":      {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9020 Chrome/108.0.0.0 Electron/22.3.26 Safari/537.36"},
//	}
//	req.Header = headers
//	resp1, err := client.Do(req)
//	if err != nil {
//		log.Fatal(err)
//	}
//	cookies := resp1.Cookies()
//	dcf := cookies[0].Value
//	sdc := cookies[1].Value
//	cfr := cookies[2].Value
//
//	return dcf, sdc, cfr
//
//}

func ContextProperties(guildid string, channelid string, id int) string {
	//properties := map[string]interface{}{
	//	"location":              "Join Guild",
	//	"location_guild_id":     guildid,
	//	"location_channel_id":   channelid,
	//	"location_channel_type": id,
	//}
	properties := fmt.Sprintf(`{"location":"Join Guild","location_guild_id":"%v","location_channel_id":"%v","location_channel_type":%v}`, guildid, channelid, id)
	//propertiesJSON, err := json.Marshal(properties)
	//if err != nil {
	//	fmt.Println("Error encoding properties to JSON:", err)
	//
	//}

	propertiesEncoded := base64.StdEncoding.EncodeToString([]byte(properties))
	//LogInfo("Succesfully Received Context Properties", "Value", Replacelast(propertiesEncoded))
	return propertiesEncoded

}
