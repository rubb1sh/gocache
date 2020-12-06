package gocache

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	gca := Init()
	gca.Run()

	gca.Add("123213", "HHH", time.Second*10)
}
