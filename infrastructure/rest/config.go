package rest

import (
	"fmt"
)

type configFn func(*config)

func WithHost(host string) configFn {
	return func(c *config) {
		c.host = host
	}
}

func WithRateLimit() configFn {
	return func(c *config) {

	}
}

type config struct {
	host string
	port int
}

func (c *config) address() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

func defaultConfig() config {
	return config{
		host: "",
	}
}
