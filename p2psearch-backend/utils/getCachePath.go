package utils

import (
	"github.com/alexkefer/p2psearch-backend/log"
	"os"
	"runtime"
)

// GetCachePath returns the file system path where cached files should be stored. Returns and logs error if the user's
// home directory couldn't be found for some reason.
func GetCachePath() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Error("couldn't get users home directory")
	}

	switch runtime.GOOS {
	case "windows":
		// Maybe on windows this should be somewhere else like appdata?
		return home + "\\p2pwebcache", err
	default:
		return home + "/.cache/p2pwebcache", err
	}
}
