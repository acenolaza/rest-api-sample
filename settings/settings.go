package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const defaultEnv = "preproduction"

var environments = map[string]string{
	"production":    "settings/prod.json",
	"preproduction": "settings/pre.json",
	"tests":         "settings/tests.json",
}

var settingsInstance *Settings

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

func (s *Settings) init() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		log.Printf("Warning: Setting %v environment due to lack of GO_ENV value", defaultEnv)
		env = defaultEnv
	}
	loadSettingsByEnv(env)
}

func loadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		log.Println("Error while reading config file", err)
	}

	jsonErr := json.Unmarshal(content, &settingsInstance)
	if jsonErr != nil {
		log.Println("Error while parsing config file", jsonErr)
	}
}

func GetSettings() *Settings {
	if settingsInstance == nil {
		settingsInstance.init()
	}
	return settingsInstance
}
