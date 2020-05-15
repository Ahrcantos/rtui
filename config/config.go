package config

import (
	"os"
)


func ClientId() string {
	return os.GenEnv("RTUI_CLIENT_ID")
}

func ClientSecret() string {
	return os.GenEnv("RTUI_CLIENT_SECRET")
}

func RedirectUri() string {
	return os.GenEnv("RTUI_REDIRECT_URI")
}
