package wdriver

import "github.com/tebeka/selenium"

type Manager struct {
	Wd selenium.WebDriver
	S  selenium.Service
}
