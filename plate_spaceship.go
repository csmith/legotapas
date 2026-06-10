//go:build lego_spaceship

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/spaceship"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "spaceship" {
		return nil, fmt.Errorf("this build of legotapas only supports `spaceship` as a provider")
	}

	return spaceship.NewDNSProvider()
}
