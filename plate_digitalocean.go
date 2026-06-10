//go:build lego_digitalocean

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/digitalocean"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "digitalocean" {
		return nil, fmt.Errorf("this build of legotapas only supports `digitalocean` as a provider")
	}

	return digitalocean.NewDNSProvider()
}
