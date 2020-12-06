package main

import (
	"os"
	"testing"
	"time"

	"github.com/rubb1sh/gocache"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPostiveDeleteKey(t *testing.T) {
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

	assert.Equal(t, gocache.Get("12"), "HHH", "The two words should be the same.")
	time.Sleep(6 * time.Second)
	assert.Equal(t, gocache.Get("12"), "", "The two words should be the same.")

}

func TestPostiveDeleteKeys(t *testing.T) {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	gocache := gocache.Init()
	gocache.Run()

	err := gocache.Add("12", "HHH", time.Second*5)
	err = gocache.Add("11", "HHH", time.Second*5)
	if err != nil {
		log.Debugf("error add")
	}

	time.Sleep(6 * time.Second)
	assert.Equal(t, gocache.Get("12"), "", "The two words should be the same.")
	assert.Equal(t, gocache.Get("11"), "", "The two words should be the same.")

}

func TestGetLen(t *testing.T) {
	gocache := gocache.Init()
	gocache.Run()

	gocache.Add("12", "HHH", time.Second*5)
	gocache.Add("11", "HHH", time.Second*5)

	assert.Equal(t, gocache.Len(), 2, "The two words should be the same.")
	time.Sleep(6 * time.Second)

	assert.Equal(t, gocache.Len(), 0, "The two words should be the same.")
}
