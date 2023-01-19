package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLog(c *ShareConfig) {
	log.SetOutput(os.Stdout)
	if c.Environment == "development" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}
