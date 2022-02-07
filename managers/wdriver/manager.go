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

	return Manager{
		S: *service,
	}
}

func (m *Manager) GetClientWebDriver(port int, proxy string, proxyPort int) selenium.WebDriver {
	caps := selenium.Capabilities{"browserName": "firefox"}
	caps.AddProxy(selenium.Proxy{
		Type:      "system",
		SOCKS:     proxy,
		SocksPort: proxyPort,
	})

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}

	return wd
}
