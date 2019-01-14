package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
)

// ChatConfig stores service config parameters
type ChatConfig struct {
	Host                 string `toml:"HOST" env:"CHATCONFIG_HOST" envDefault:":8080"`
	Debug                bool   `toml:"DEBUG" env:"CHATCONFIG_DEBUG" envDefault:"false"`
	Noauth               bool   `toml:"NOAUTH" env:"CHATCONFIG_NOAUTH" envDefault:"false"`
	FacebookClientID     string `toml:"FACEBOOK_CLIENT_ID" env:"CHATCONFIG_FACEBOOK_CLIENT_ID,required"`
	FacebookClientSecret string `toml:"FACEBOOK_CLIENT_SECRET" env:"CHATCONFIG_FACEBOOK_CLIENT_SECRET,required"`
	GithubClientID       string `toml:"GITHUB_CLIENT_ID" env:"CHATCONFIG_GITHUB_CLIENT_ID,required"`
	GithubClientSecret   string `toml:"GITHUB_CLIENT_SECRET" env:"CHATCONFIG_GITHUB_CLIENT_SECRET,required"`
	GoogleClientID       string `toml:"GOOGLE_CLIENT_ID" env:"CHATCONFIG_GOOGLE_CLIENT_ID,required"`
	GoogleClientSecret   string `toml:"GOOGLE_CLIENT_SECRET" env:"CHATCONFIG_GOOGLE_CLIENT_SECRET,required"`
}

// Help output for flags when program run with -h flag
/*
func setFlagsHelp() map[string]string {
	usageMsg := make(map[string]string)

	usageMsg["host"] = "server host address"

	usageMsg["Debug"] = "debug mode enables tracing of events"
	usageMsg["noauth"] = "allow to use chat without authentication"
	usageMsg["FacebookClientID"] = "Facebook client id for OAUTH"
	usageMsg["FacebookClientSecret"] = `Facebook client secret for OAUTH`
	usageMsg["GithubClientID"] = "Github client id for OAUTH"
	usageMsg["GithubClientSecret"] = `Github client secret for OAUTH`
	usageMsg["GoogleClientID"] = "Google client id for OAUTH"
	usageMsg["GoogleClientSecret"] = `Google client secret for OAUTH`

	return usageMsg
}
*/

// Load loads configuration for service to struct ChatConfig
func Load(configPath string) *ChatConfig {
	svcConfig := &ChatConfig{}

	log.Printf("Loading configuration\n")
	if configPath != "" {
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			log.Printf("Provided local config file [%s] does not exist. Flags and Environment variables will be used ", configPath)
			if err = env.Parse(svcConfig); err != nil {
				log.Fatalf("Load: failed to parse env: %v", err)

			}

		} else {
			log.Printf("Local config file [%s] will be used", configPath)
			// Choose what while is passed
			if strings.HasSuffix(configPath, "toml") {
				log.Printf("Toml detected")
				if _, err = toml.DecodeFile("config.toml", svcConfig); err != nil {
					log.Fatalf("Load: failed to parse toml%v", err)
				}
			}

		}
	}
	prettyConfig, err := json.MarshalIndent(svcConfig, "", "")
	if err != nil {
		log.Fatalf("Failed to marshal indent config: %v", err)

	}
	log.Printf("Current config:\n %s", string(prettyConfig))

	return svcConfig

}
