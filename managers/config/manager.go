package config

import (
	"encoding/json"
	"github.com/Vladimir-Urik/AutoVote/managers/craftlist"
	"github.com/Vladimir-Urik/AutoVote/managers/czechcraft"
	"os"
)

func LoadConfigFromFile(file string) Config {
	var config Config
	if !fileExists(file) {
		config = createConfig(file)
	} else {
		config = loadConfig(file)
	}
	return config
}

func loadConfig(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}

func createConfig(file string) Config {
	config := Config{
		CaptchaSettings: &CaptchaSettings{
			Key: "-",
		},
		CzechCraftSettings: &czechcraft.Settings{
			Name:    "-",
			Path:    "-",
			SiteKey: "-",
		},
		CraftListSettings: &craftlist.Settings{
			Name:    "-",
			Path:    "-",
			SiteKey: "-",
		},
	}
	saveConfig(file, config)
	return config
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func saveConfig(file string, cfg Config) {
	configFile, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			panic(err)
			return
		}
	}(configFile)

	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(cfg)
	if err != nil {
		panic(err)
	}
}
