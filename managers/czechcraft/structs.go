package czechcraft

import (
	"github.com/Vladimir-Urik/AutoVote/managers/captcha"
	"github.com/Vladimir-Urik/AutoVote/managers/wdriver"
)

type Manager struct {
	Settings      *Settings
	CaptchaSolver *captcha.Manager
	WebDriver     *wdriver.Manager
}

type Settings struct {
	Name    string
	Path    string
	SiteKey string
}
