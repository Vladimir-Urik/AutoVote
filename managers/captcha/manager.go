package captcha

import (
	api2captcha "github.com/2captcha/2captcha-go"
	"github.com/Vladimir-Urik/AutoVote/logger"
)

func LoadCaptchaSolver(key string) Manager {
	logger.Info("Loading captcha solver...")
	return Manager{client: api2captcha.NewClient(key)}
}

func (m *Manager) Solve(settings api2captcha.ReCaptcha) (string, error) {
	return m.client.Solve(settings.ToRequest())
}
