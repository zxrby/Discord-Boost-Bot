package Discord

import (
	tls_client "github.com/bogdanfinn/tls-client"
)

type Discord struct {
	Client            tls_client.HttpClient
	Token             string
	SubscriptionSlots SubscriptionSlots
	Fingerprint       string
	Proxy             string
	GuildId           string
	ChannelId         string
	SuperProperties   string
	ContextProperties string
	Host              string
	Auth              string
}

type SubscriptionSlots []struct {
	Id                       string      `json:"id"`
	SubscriptionId           string      `json:"subscription_id"`
	PremiumGuildSubscription interface{} `json:"premium_guild_subscription"`
	Canceled                 bool        `json:"canceled"`
	CooldownEndsAt           interface{} `json:"cooldown_ends_at"`
}

type ServerJoinRQ struct {
	CaptchaKey     []string `json:"captcha_key"`
	CaptchaSitekey string   `json:"captcha_sitekey"`
	CaptchaService string   `json:"captcha_service"`
	CaptchaRqdata  string   `json:"captcha_rqdata"`
	CaptchaRqtoken string   `json:"captcha_rqtoken"`
}

type GuildInfo struct {
	Type      int    `json:"type"`
	Code      string `json:"code"`
	ExpiresAt any    `json:"expires_at"`
	Guild     struct {
		ID                       string   `json:"id"`
		Name                     string   `json:"name"`
		Splash                   any      `json:"splash"`
		Banner                   any      `json:"banner"`
		Description              any      `json:"description"`
		Icon                     any      `json:"icon"`
		Features                 []string `json:"features"`
		VerificationLevel        int      `json:"verification_level"`
		VanityURLCode            string   `json:"vanity_url_code"`
		NsfwLevel                int      `json:"nsfw_level"`
		Nsfw                     bool     `json:"nsfw"`
		PremiumSubscriptionCount int      `json:"premium_subscription_count"`
	} `json:"guild"`
	GuildID string `json:"guild_id"`
	Channel struct {
		ID   string `json:"id"`
		Type int    `json:"type"`
		Name string `json:"name"`
	} `json:"channel"`
	ApproximateMemberCount   int `json:"approximate_member_count"`
	ApproximatePresenceCount int `json:"approximate_presence_count"`
}
