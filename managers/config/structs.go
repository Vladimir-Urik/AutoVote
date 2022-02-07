package config

import (
	"github.com/Vladimir-Urik/AutoVote/managers/craftlist"
	"github.com/Vladimir-Urik/AutoVote/managers/czechcraft"
)

type Config struct {
	CaptchaSettings    *CaptchaSettings
	CzechCraftSettings *czechcraft.Settings
	CraftListSettings  *craftlist.Settings
	VoteWebhook        string
	LogsWebhook        string
	Proxies            []string
}

type CaptchaSettings struct {
	Key string
}

type WebhookSettings struct {
	URL     string
	Enabled bool
	Format  []string
}
