package craftlist

import (
	"github.com/Vladimir-Urik/AutoVote/managers/captcha"
	"github.com/Vladimir-Urik/AutoVote/managers/config"
	"github.com/Vladimir-Urik/AutoVote/managers/wdriver"
)

type Manager struct {
	Config        *config.Config
	CaptchaSolver *captcha.Manager
	WebDriver     *wdriver.Manager
}

type Settings struct {
	Name    string
	Path    string
	SiteKey string
}
