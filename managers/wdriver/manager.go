package wdriver

import (
	"fmt"
	"github.com/Vladimir-Urik/AutoVote/logger"
	"github.com/tebeka/selenium"
	"os"
)

func CreateNewWDriver(port int) Manager {
	const (
		seleniumPath    = "./vendor/selenium-server.jar"
		geckoDriverPath = "./vendor/geckodriver"
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),
		selenium.GeckoDriver(geckoDriverPath),
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(false)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		logger.Info(fmt.Sprintf("Error starting selenium service: %s", err))
	}

	/*defer func(service *selenium.Service) {
		err := service.Stop()
		logger.Info(fmt.Sprintf("Stopped Selenium service: %v", err))
		if err != nil {
			panic(err)
		}
	}(service)*/

	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}

	/*defer func(wd selenium.WebDriver) {
		err := wd.Quit()
		logger.Info(fmt.Sprintf("Quit WebDriver: %v", err))
		if err != nil {
			panic(err)
		}
	}(wd)*/

	return Manager{
		Wd: wd,
		S:  *service,
	}
}
