package loadenv

import "os"

type Config struct {
	ApiBffPort        string
	ApiBackGetList    string
	ApiBackGetSearch  string
	ApiBackPostCreate string
	ApiBackPatchEdit  string
}

func LoadEnvFile() *Config {

	return &Config{
		ApiBffPort:        os.Getenv("API_BFF_PORT"),
		ApiBackGetList:    os.Getenv("API_BACK_GET_LIST"),
		ApiBackGetSearch:  os.Getenv("API_BACK_GET_SEARCH"),
		ApiBackPostCreate: os.Getenv("API_BACK_POST_CREATE"),
		ApiBackPatchEdit:  os.Getenv("API_BACK_PATCH_EDIT"),
	}
}
