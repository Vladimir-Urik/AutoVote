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

	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			panic(err)
		}
	}(configFile)
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
		VoteWebhook: "-",
		LogsWebhook: "-",
		Proxies: []string{
			/* Czech Socks4 */
			"89.203.220.110:4153",
			"109.238.219.241:4153",
			"78.102.14.1:5678",
			"81.19.3.249:10080",
			"185.175.221.65:4153",
			"88.146.204.49:4153",
			"93.91.146.30:34350",
			"217.197.145.198:5678",
			"185.70.218.13:42707",
			"89.190.44.185:4153",
			"81.162.199.66:4153",
			"46.174.56.21:5678",
			"194.228.84.10:4145",
			"89.203.129.202:4153",
			"109.238.222.1:4153",
			"176.98.248.2:4153",
			"31.7.243.190:1080",
			"188.175.207.27:5678",
			"188.75.186.152:4145",
			"91.187.63.236:5678",
			"185.32.181.68:4153",
			"109.183.189.238:36729",
			"85.135.95.218:4145",
			"109.238.208.138:51372",
			"77.48.137.3:50523",

			/* Slovak Socks4 */
			"82.119.98.122:4153",
			"185.29.156.233:10801",
			"217.145.199.47:56746",
			"185.18.64.84:4153",
			"80.81.232.145:5678",
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
