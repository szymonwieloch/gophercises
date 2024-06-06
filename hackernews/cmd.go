package main

import (
	"fmt"
	"time"
)

type cacheChoice int

const (
	useCacheNone cacheChoice = iota
	useCacheRefres
	useCacheBackground
)

func (e *cacheChoice) UnmarshalText(b []byte) error {
	switch string(b) {
	case "none":
		*e = useCacheNone
	case "refresh":
		*e = useCacheRefres
	case "background":
		*e = useCacheBackground
	default:
		return fmt.Errorf("invalid cache type: %q: options are 'none', 'refresh', 'background'", string(b))
	}
	return nil
}

type args struct {
	Cache  cacheChoice   `arg:"--cache" help:"Kind of cache to use. Options: 'none', 'refresh', 'background'" default:"none"`
	Count  uint16        `arg:"-e,--entries" help:"Number of entries shown on the main page" default:"30"`
	Port   uint16        `arg:"--port" help:"Port to launch the server on" default:"3000"`
	Period time.Duration `arg:"-p,--period" help:"Sets period for cache refresh" default:"30s"`
}
