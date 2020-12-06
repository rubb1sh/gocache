package gocache

import (
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type timeStamp int64

var (
	TTL_ERROR = fmt.Errorf("input time is wrong, need to be mod by time.Second and more than zero.")
)

type GoCache struct {
	_Map    map[string]string
	timeMap map[timeStamp][]string
	lock    *sync.RWMutex
}

func Init() *GoCache {
	return &GoCache{
		_Map:    make(map[string]string),
		timeMap: make(map[timeStamp][]string),
		lock:    new(sync.RWMutex),
	}
}

func (g *GoCache) Add(k, v string, ttl time.Duration) error {

	if ttl < 0 || ttl%time.Second != 0 {
		return TTL_ERROR
	}

	g.lock.Lock()
	defer g.lock.Unlock()
	g._Map[k] = v
	t := time.Now().Add(ttl)
	_, ok := g.timeMap[timeStamp(t.Unix())]
	if !ok {
		g.timeMap[timeStamp(t.Unix())] = make([]string, 0)
	}
	g.timeMap[timeStamp(t.Unix())] = append(g.timeMap[timeStamp(t.Unix())], k)

	log.Debugf("add a key, k=%s, v=%s, ttl=%v", k, v, t)
	return nil
}

func (g *GoCache) Run() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for t := range ticker.C {
			if _, ok := g.timeMap[timeStamp(t.Unix())]; ok {
				g.deleteWithTTL(timeStamp(t.Unix()))
			}
		}
		ticker.Stop()

	}()

}

func (g *GoCache) deleteWithTTL(ts timeStamp) {
	g.lock.Lock()
	defer g.lock.Unlock()
	keys, ok := g.timeMap[ts]
	if !ok {
		return
	}
	log.Debugf("delete key, key=%+v, now=%v", keys, time.Unix(int64(ts), 0))
	for _, key := range keys {
		delete(g._Map, key)
		log.Debugf("delete key, key=%s, now=%v", key, time.Unix(int64(ts), 0))
	}
	delete(g.timeMap, ts)

}

func (g *GoCache) Get(k string) string {
	g.lock.RLock()
	defer g.lock.RUnlock()

	log.Debugf("get key, key=%s, value=%s", k, g._Map[k])
	return g._Map[k]
}

func (g *GoCache) Len() int {
	log.Debugf("get map length, length=%d", len(g._Map))
	return len(g._Map)
}

func (g *GoCache) Delete(k string) {
	g.lock.Lock()
	defer g.lock.Unlock()
	_, ok := g._Map[k]
	if !ok {
		return
	}
	//delete(g.timeMap, ts) timeMap is exist, postive delete can't delete timeMap key
	delete(g._Map, k)
	log.Debugf("delete key, key=%s", k)
}
