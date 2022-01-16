package config

import (
	"github.com/Vladimir-Urik/AutoVote/managers/craftlist"
	"github.com/Vladimir-Urik/AutoVote/managers/czechcraft"
)

type Config struct {
	CaptchaSettings    *CaptchaSettings
	CzechCraftSettings *czechcraft.Settings
	CraftListSettings  *craftlist.Settings
}

type CaptchaSettings struct {
	Key string
}
