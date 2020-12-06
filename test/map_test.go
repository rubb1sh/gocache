package main

import (
	"os"
	"testing"
	"time"

	"github.com/rubb1sh/gocache"
	log "github.com/sirupsen/logrus"
)

func TestHelloWorld(t *testing.T) {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	gocache := gocache.Init()
	gocache.Run()

	err := gocache.Add("12", "HHH", time.Second*5)
	if err != nil {
		log.Debugf("error add")
	}

	log.Debugf("ggg %s", gocache.Get("12"))
	time.Sleep(6 * time.Second)
	log.Debugf("ggg %s", gocache.Get("12"))

	for {

	}
}
