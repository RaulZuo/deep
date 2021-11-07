package options

import (
	"github.com/RaulZuo/deep/internal/pkg/server"
	"net"
	"strconv"
)

type InsecureServingOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
}

// NewInsecureServingOptions is for creating an unauthenticated, unauthorized, insecure port.
func NewInsecureServingOptions() *InsecureServingOptions {
	return &InsecureServingOptions{
		BindAddress: "127.0.0.1",
		BindPort:    8080,
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (s *InsecureServingOptions) ApplyTo(c *server.Config) error {
	c.InsecureServing = &server.InsecureServingInfo{
		Address: net.JoinHostPort(s.BindAddress, strconv.Itoa(s.BindPort)),
	}

	return nil
}
