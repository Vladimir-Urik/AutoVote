package main

import (
	"github.com/Vladimir-Urik/AutoVote/logger"
	"github.com/Vladimir-Urik/AutoVote/managers/captcha"
	"github.com/Vladimir-Urik/AutoVote/managers/config"
	"github.com/Vladimir-Urik/AutoVote/managers/craftlist"
	"github.com/Vladimir-Urik/AutoVote/managers/czechcraft"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Version = "0.0.1"

func main() {
	start := time.Now().UnixNano() / int64(time.Millisecond)

	logger.SetupLogging()
	logger.Info("Booting...")
	logger.Info(" ")
	logger.Info("------------------------------")
	logger.Info(" Developer: GGGEDR")
	logger.Info(" Version: " + Version)
	logger.Info("------------------------------")
	logger.Info(" ")
	logger.Info("Starting AutoVoting system...")

	logger.Info("Loading config...")
	cfg := config.LoadConfigFromFile("config.json")
	logger.Info("Config loaded")

	logger.Info("Starting captcha manager...")
	captchaSolver := captcha.LoadCaptchaSolver(cfg.CaptchaSettings.Key)
	logger.Info("Captcha manager started")

	logger.Info("Starting CzechCraft manager...")
	cc := czechcraft.StartCzechCraftManager(cfg.CzechCraftSettings, &captchaSolver)
	logger.Debug("Calling CzechCraft Thread")
	cc.StartVotingThread()
	logger.Info("CzechCraft manager started")

	logger.Info("Starting CraftList manager...")
	cl := craftlist.StartCraftListManager(cfg.CraftListSettings, &captchaSolver)
	logger.Debug("Calling CraftList Thread")
	cl.StartVotingThread()
	logger.Info("CraftList manager started")

	end := time.Now().UnixNano() / int64(time.Millisecond)
	difference := end - start

	logger.InfoFm("System started in ", difference, " ms\n")

	done := make(chan struct{})

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		close(done)
	}()

	<-done
	logger.Info("Shutting down...")
	err := cl.WebDriver.Wd.Quit()
	if err != nil {
		return
	}

	err = cl.WebDriver.S.Stop()
	if err != nil {
		return
	}

	err = cc.WebDriver.Wd.Quit()
	if err != nil {
		return
	}

	err = cc.WebDriver.S.Stop()
	if err != nil {
		return
	}

	logger.Info("Shutdown complete")
}
