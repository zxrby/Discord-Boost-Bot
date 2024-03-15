package Discord

import (
	"BoostTool/Core/Captcha"
	"BoostTool/Core/Utils"
	"encoding/json"
	"errors"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/charmbracelet/log"
	"io"
	"math/rand"
	"net/url"
	"strings"
)

func New(token string) Discord {
	jar := tls_client.NewCookieJar([]tls_client.CookieJarOption{}...)
	proxy := Utils.Proxy()
	host := strings.Split(proxy, "@")[1]
	auth := strings.Split(proxy, "@")[0]

	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(config.Timeout),
		tls_client.WithClientProfile(profiles.Chrome_107),
		tls_client.WithCookieJar(jar),
	}
	if config.Proxyless == false {
		options = append(options, tls_client.WithProxyUrl("http://"+proxy))
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		Utils.LogError(err.Error(), "", "")
	}
	client.SetCookies(&url.URL{Path: "/", Host: "discord.com", Scheme: "https"}, []*http.Cookie{{Name: "locale", Value: "en-US"}})

	c := Discord{
		Token:  Utils.FormatToken(token),
		Client: client,
		Proxy:  proxy,
		Host:   host,
		Auth:   auth,
	}

	return c
}

func (c *Discord) GetHeaders() http.Header {

	headers := http.Header{
		"authority":          {"discord.com"},
		"accept":             {"*/*"},
		"accept-language":    {"fr-FR,fr;q=0.9,en-US,pa;q=0.9"},
		"authorization":      {c.Token},
		"cache-control":      {"no-cache"},
		"content-type":       {"application/json"},
		"origin":             {"https://discord.com"},
		"pragma":             {"no-cache"},
		"referer":            {"https://discord.com/channels/@me"},
		"sec-ch-ua":          {`"Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`},
		"sec-ch-ua-mobile":   {"?0"},
		"sec-ch-ua-platform": {`"Windows"`},
		"sec-fetch-dest":     {"empty"},
		"sec-fetch-mode":     {"cors"},
		"sec-fetch-site":     {"same-origin"},
		"user-agent":         {"Mozilla/5.0 (Macintosh; U; PPC Mac OS X; de-de) AppleWebKit/85.8.5 (KHTML, like Gecko) Safari/85"},
		"x-debug-options":    {"bugReporterEnabled"},
		"x-discord-locale":   {"en-US"},
		"x-super-properties": {"eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IlNhZmFyaSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJlbi1KTSIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IFU7IFBQQyBNYWMgT1MgWDsgZGUtZGUpIEFwcGxlV2ViS2l0Lzg1LjguNSAoS0hUTUwsIGxpa2UgR2Vja28pIFNhZmFyaS84NSIsImJyb3dzZXJfdmVyc2lvbiI6IiIsIm9zX3ZlcnNpb24iOiIiLCJyZWZlcnJlciI6IiIsInJlZmVycmluZ19kb21haW4iOiIiLCJyZWZlcnJlcl9jdXJyZW50IjoiIiwicmVmZXJyaW5nX2RvbWFpbl9jdXJyZW50IjoiIiwicmVsZWFzZV9jaGFubmVsIjoic3RhYmxlIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTgxODMyLCJjbGllbnRfZXZlbnRfc291cmNlIjoibnVsbCJ9"},
	}

	return headers

}

func (c *Discord) GetRequiredElements() error {
    client := c.Client
    req, err := http.NewRequest("GET", "https://discord.com/api/v9/experiments?with_guild_experiments=true", nil)
    if err != nil {
        Utils.LogError("Failed to create HTTP request", "", "")
        return err
    }

    resp, err := client.Do(req)
    if err != nil {
        Utils.LogError("Failed to send HTTP request", "", "")
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        Utils.LogError("Non-OK status code received", "", "")
        return errors.New("Non-OK status code received")
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        Utils.LogError("Failed to read response body", "", "")
        return err
    }

    var reply FingerprintResponse
    if err := json.Unmarshal(body, &reply); err != nil {
        Utils.LogError("Failed to unmarshal JSON response", "", "")
        return err
    }

    c.Fingerprint = reply.Fingerprint
    return nil
}


func (c *Discord) IsValidInvite(invite string) error {

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://discord.com/api/v9/invites/%v", invite), nil)
	resp, err := c.Client.Do(req)
	body, _ := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Invalid Invite")
	}
	var reply GuildInfo
	if err = json.Unmarshal(body, &reply); err != nil {
		return err
	}
	c.GuildId = reply.Guild.ID
	c.ChannelId = reply.Channel.ID
	c.ContextProperties = Utils.ContextProperties(reply.Guild.ID, reply.Channel.ID, reply.Channel.Type)
	return err
}

func (c *Discord) JoinServer(invite string) error {
	var FormattedToken string
	FormattedToken = Utils.Replacelast(c.Token)
	headers := c.GetHeaders()
	headers.Set("x-context-properties", c.ContextProperties)
	var data = strings.NewReader(`{}`)
	req, err := http.NewRequest("POST", "https://discord.com/api/v9/invites/"+invite, data)
	if err != nil {
		log.Error(err.Error())
		return errors.New("Failed Joining Guild")
	}
	req.Header = headers
	req.Header.Del("Content-Length")
	resp1, err := c.Client.Do(req)
	if err != nil {
		log.Error(err.Error())
		return errors.New("Failed Joining Guild")
	}
	defer resp1.Body.Close()

	bodytext, err := io.ReadAll(resp1.Body)
	if err != nil {
		log.Error(err.Error())
		return errors.New("Failed Joining Guild")
	}

	if resp1.StatusCode == 200 {
		Utils.LogSuccess("Successfully Joined Server: discord.gg/"+invite, "Token", FormattedToken)
		return nil
	} else if strings.Contains(string(bodytext), "captcha_key") {
		var solution string
		var rqdata ServerJoinRQ

		Utils.LogInfo("Encountered Captcha", "token", FormattedToken)
		err = json.Unmarshal(bodytext, &rqdata)
		if err != nil {
			log.Error(err.Error())
			return errors.New("Failed Joining Guild")
		}

		var captchaServiceFunc func(string, string, string, string, string, string) string

		switch strings.ToLower(config.CapService) {
		case "capsolver":
			captchaServiceFunc = Captcha.Captcha
		case "hcoptcha":
			captchaServiceFunc = Captcha.Hcoptcha
		case "capmonster":
			captchaServiceFunc = Captcha.Capmonster
		default:
			log.Error("Invalid captcha service configuration")
			return errors.New("Failed Joining Guild")
		}

		for i := 0; i < 3; i++ {
			solution = captchaServiceFunc(config.CapKey, c.Host, c.Auth, rqdata.CaptchaSitekey, rqdata.CaptchaRqdata, c.Token)
			if solution != "" {
				break
			}
		}

		payload := fmt.Sprintf(`{"captcha_key": "%v", "captcha_rqtoken": "%v"}`, solution, rqdata.CaptchaRqtoken)

		req2, err := http.NewRequest("POST", "https://discord.com/api/v9/invites/"+invite, strings.NewReader(payload))
		if err != nil {
			log.Error(err.Error())
			return errors.New("Failed Joining Guild")
		}
		req2.Header = headers
		req2.Header.Del("Content-Length")

		resp2, err := c.Client.Do(req2)
		if err != nil {
			log.Error(err.Error())
			return errors.New("Failed Joining Guild")
		}
		defer resp2.Body.Close()

		bodytext2, err := io.ReadAll(resp2.Body)
		if err != nil {
			log.Error(err.Error())
			return errors.New("Failed Joining Guild")
		}

		if resp2.StatusCode == 200 {
			Utils.LogInfo("Captcha Solved", "Token", FormattedToken)
			Utils.LogSuccess("Successfully Joined Server: discord.gg/"+invite, "Token", FormattedToken)
			return nil
		} else if strings.Contains(string(bodytext2), "Verify") {
			Utils.LogError("Token is Locked", "Token", FormattedToken)
			return errors.New("Failed Joining Guild")
		} else if resp1.StatusCode == 401 {
			Utils.LogError("Invalid Token", "Token", FormattedToken)
			return errors.New("Failed Joining Guild")
		} else {
			Utils.LogError("Received Unknown Error: "+string(bodytext2), "Token", FormattedToken)
			return errors.New("Failed Joining Guild")
		}
	} else if strings.Contains(string(bodytext), "Verify") {
		Utils.LogError("Token is Locked", "Token", FormattedToken)
		return errors.New("Failed Joining Guild")
	} else if resp1.StatusCode == 401 {
		Utils.LogError("Invalid Token", "Token", FormattedToken)
		return errors.New("Failed Joining Guild")
	} else {
		Utils.LogError("Received Unknown Response Error"+string(bodytext), "Token", FormattedToken)
		return errors.New("Failed Joining Guild")
	}
}


func (c *Discord) GetSubscriptionSlots() error {
	var FormattedToken string
	go func() {
		FormattedToken = Utils.Replacelast(c.Token)
	}()
	headers := c.GetHeaders()
	headers["referer"] = []string{"https://discord.com/channels/" + c.GuildId + "/" + c.ChannelId}

	req, err := http.NewRequest(http.MethodGet, "https://discord.com/api/v9/users/@me/guilds/premium/subscription-slots", nil)
	req.Header = headers
	req.Header.Del("Content-Length")

	resp, err := c.Client.Do(req)

	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var reply SubscriptionSlots
	if err = json.Unmarshal(body, &reply); err != nil {
		return err
	}
	if len(reply) == 0 {
		Utils.LogError("No Boosts Available, Token has been Used or Nitro Revoked!", "Token", FormattedToken)
		err1 := errors.New("No Boosts Available, Token has been Used or Nitro Revoked!")
		return err1
	} else {
		c.SubscriptionSlots = reply
	}
	return nil
}

func (c *Discord) BoostServer() error {
	var FormattedToken string
	go func() {
		FormattedToken = Utils.Replacelast(c.Token)
	}()
	headers := c.GetHeaders()
	path := fmt.Sprintf("/api/v9/guilds/%v/premium/subscriptions", c.GuildId)

	var payload = strings.NewReader(`{"user_premium_guild_subscription_slot_ids":["` + c.SubscriptionSlots[0].Id + `","` + c.SubscriptionSlots[1].Id + `"]}`)

	req, err := http.NewRequest(http.MethodPut, "https://discord.com/api/v9/guilds/"+c.GuildId+"/premium/subscriptions", payload)
	req.Header = headers
	req.Header.Add("path", path)

	resp, err := c.Client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return errors.New("Failed to Boost!")
	} else {
		Utils.LogSuccess("Successfully Boosted Server", "Token", FormattedToken)
	}

	return nil
}

func (c *Discord) CustomizeTokens() {
	var needed = 0
	var customized = 0
	var token1 string

	go func() {
		token1 = Utils.Replacelast(c.Token)
	}()

	headers := c.GetHeaders()
	if config.CustomPersonalization.DisplayName != "" {
		needed += 1

		var data = strings.NewReader(`{"nick":"` + config.CustomPersonalization.DisplayName + `"}`)
		req, _ := http.NewRequest(http.MethodPatch, "https://discord.com/api/v9/guilds/"+c.GuildId+"/members/@me", data)
		req.Header = headers

		resp1, err := c.Client.Do(req)
		defer resp1.Body.Close()

		if err != nil {
			log.Error(err)
		}

		bodyText, err := io.ReadAll(resp1.Body)

		if resp1.StatusCode == 200 {
			customized += 1
		} else {
			Utils.LogError("Failed Changing Display Name", "Error", string(bodyText))
		}

	}

	if config.CustomPersonalization.CustomBio != "" {
		needed += 1
		var data = fmt.Sprintf(`{"bio": "%v"}`, config.CustomPersonalization.CustomBio)

		req, _ := http.NewRequest(http.MethodPatch, "https://discord.com/api/v9/users/@me/profile", strings.NewReader(data))
		req.Header = headers
		req.Header.Del("Content-Length")

		resp1, err := c.Client.Do(req)
		defer resp1.Body.Close()

		if err != nil {
			log.Error(err)
		}

		bodyText, err := io.ReadAll(resp1.Body)
		if resp1.StatusCode == 200 {
			customized += 1
		} else {
			Utils.LogError("Failed Changing Bio", "Error", string(bodyText))
		}

	}

	if len(config.CustomPersonalization.CustomPfp) != 0 {
		pfp := config.CustomPersonalization.CustomPfp[rand.Intn(len(config.CustomPersonalization.CustomPfp))]
		needed += 1
		bs4 := Utils.ImageToB64("./Data/Avatars/" + pfp)
		var data = strings.NewReader(`{"avatar":"` + bs4 + `"}`)
		req, _ := http.NewRequest(http.MethodPatch, "https://discord.com/api/v9/guilds/"+c.GuildId+"/members/@me", data)
		req.Header = headers
		req.Header.Del("Content-Length")

		resp1, err := c.Client.Do(req)
		defer resp1.Body.Close()

		if err != nil {
			log.Error(err)
		}
		bodyText, err := io.ReadAll(resp1.Body)
		if resp1.StatusCode == 200 {
			customized += 1
		} else {
			Utils.LogError("Failed Changing Avatar", "Error", string(bodyText))
		}

	}

	if len(config.CustomPersonalization.CustomBanner) != 0 {
		banner := config.CustomPersonalization.CustomBanner[rand.Intn(len(config.CustomPersonalization.CustomBanner))]
		needed += 1
		bs4 := Utils.ImageToB64("./Data/Banners/" + banner)
		var data = strings.NewReader(`{"banner":"` + bs4 + `"}`)
		req, _ := http.NewRequest(http.MethodPatch, "https://discord.com/api/v9/guilds/"+c.GuildId+"/profile/@me", data)
		req.Header = headers
		req.Header.Del("Content-Length")

		resp1, err := c.Client.Do(req)
		defer resp1.Body.Close()

		if err != nil {
			log.Error(err)
		}
		bodyText, err := io.ReadAll(resp1.Body)
		if resp1.StatusCode == 200 {
			customized += 1
		} else {
			Utils.LogError("Failed Changing Banner", "Error", string(bodyText))
		}

	}

	sprint := fmt.Sprintf("Successfully Watermarked Token %v/%v Times", customized, needed)
	Utils.LogSuccess(sprint, "Token", token1)
	return
}
