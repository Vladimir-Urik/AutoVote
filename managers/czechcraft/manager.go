package czechcraft

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"github.com/Vladimir-Urik/AutoVote/logger"
	"github.com/Vladimir-Urik/AutoVote/managers/captcha"
	"github.com/Vladimir-Urik/AutoVote/managers/config"
	"github.com/Vladimir-Urik/AutoVote/managers/wdriver"
	"github.com/Vladimir-Urik/AutoVote/managers/webhook"
	"github.com/tebeka/selenium"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func StartCzechCraftManager(config *config.Config, captcha *captcha.Manager) Manager {
	logger.Info("Starting CzechCraft WebDriver...")
	wd := wdriver.CreateNewWDriver(9800)
	logger.Info("CzechCraft WebDriver started")
	return Manager{
		Config:        config,
		CaptchaSolver: captcha,
		WebDriver:     &wd,
	}
}

func (m *Manager) StartVotingThread() {
	logger.Info("Starting CzechCraft voting thread...")
	go func() {
		for {
			logger.Info("CzechCraft: Starting vote process...")
			m.vote()
			logger.Info("CzechCraft: Vote process finished. Sleeping...")
			m.sleep()
		}
	}()
	logger.Info("CzechCraft voting thread started")
}

func (m *Manager) vote() {
	var settings = m.Config.CzechCraftSettings
	logger.Info("CzechCraft: Solving captcha...")
	code, err := m.CaptchaSolver.Solve(api2captcha.ReCaptcha{
		SiteKey:   settings.SiteKey,
		Url:       "https://czech-craft.eu/server/" + settings.Path + "/vote/",
		Invisible: false,
		Action:    "verify",
	})

	if err != nil {
		logger.Error("Error while solving captcha: " + err.Error())
		return
	}

	if code == "" {
		logger.Error("Captcha code is empty")
		return
	}

	logger.Info("CzechCraft: Captcha solved: " + code)

	var proxy = m.Config.Proxies[rand.Intn(len(m.Config.Proxies))]
	if strings.Contains(proxy, ":") {
		logger.Error("Proxy is not valid: " + proxy)
		return
	}

	var proxyIp = strings.Split(proxy, ":")[0]
	var proxyPort, _ = strconv.Atoi(strings.Split(proxy, ":")[1])

	wd := m.WebDriver.GetClientWebDriver(9800, proxyIp, proxyPort)
	defer func(wd selenium.WebDriver) {
		err := wd.Quit()
		logger.Info(fmt.Sprintf("Quit WebDriver: %v", err))
		if err != nil {
			panic(err)
		}
	}(wd)

	logger.Info("CzechCraft: Opening vote page...")
	if err := wd.Get("https://czech-craft.eu/server/" + settings.Path + "/vote/"); err != nil {
		logger.Error("Error while getting page: " + err.Error())
		return
	}

	elem, err := wd.FindElement(selenium.ByID, "username")
	if err != nil {
		logger.Error("Error while finding username field: " + err.Error())
		return
	}

	err = elem.Clear()
	if err != nil {
		logger.Error("Error while clearing username field: " + err.Error())
		return
	}

	logger.Info("CzechCraft: Filling username field...")
	err = elem.SendKeys(settings.Name)
	if err != nil {
		logger.Error("Error while sending username: " + err.Error())
		return
	}

	elem, err = wd.FindElement(selenium.ByID, "privacy")
	if err != nil {
		logger.Error("Error while finding privacy field: " + err.Error())
		return
	}

	logger.Info("CzechCraft: Filling privacy field...")
	err = elem.Click()
	if err != nil {
		logger.Error("Error while clicking privacy field: " + err.Error())
		return
	}

	logger.Info("CzechCraft: Filling captcha field...")
	_, err = wd.ExecuteScript("var element=document.getElementById('g-recaptcha-response'); element.style.display='';", nil)
	if err != nil {
		logger.Error("Error while showing captcha field: " + err.Error())
		return
	}

	_, err = wd.ExecuteScript("document.getElementById('g-recaptcha-response').innerHTML = '"+code+"'", nil)
	if err != nil {
		logger.Error("Error while sending captcha code: " + err.Error())
		return
	}

	var elems []selenium.WebElement
	elems, err = wd.FindElements(selenium.ByCSSSelector, "button")
	if err != nil {
		logger.Error("Error while finding submit button: " + err.Error())
		return
	}

	if len(elems) == 0 {
		logger.Error("Submit button not found")
		return
	}

	var submitButton selenium.WebElement
	for _, elem := range elems {
		t, e := elem.Text()
		if e != nil {
			continue
		}

		if t == "Hlasovat" {
			submitButton = elem
			break
		}
	}

	if submitButton == nil {
		logger.Error("Submit button not found")
		return
	}

	logger.Info("CraftList: Submitting vote...")
	err = submitButton.Click()

	if err != nil {
		logger.Error("Error while submitting vote: " + err.Error())
		return
	}

	if err != nil {
		logger.Error("Error while quitting webdriver: " + err.Error())
		return
	}

	logger.Info("Vote is successful! Nickname: " + settings.Name + "; Path: " + settings.Path + "; Page: CzechCraft")
	var intColor = 5814783
	var embeds = []webhook.Embed{
		{
			Title:       "Úspešné hlasovanie",
			Description: "Nick: `" + settings.Name + "`\nPath: `" + settings.Path + "`\nPage: `Czech-Craft`",
			Color:       intColor,
		},
	}
	webhook.SendDataToWebhook("", embeds, m.Config.VoteWebhook)
}

func (m *Manager) sleep() {
	var seconds = rand.Intn(600-60) + 60
	var randomTime = time.Duration(seconds) * time.Second
	time.Sleep((2 * time.Hour) + randomTime)
}
