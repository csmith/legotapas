//go:build lego_lightsail

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/lightsail"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "lightsail" {
		return nil, fmt.Errorf("this build of legotapas only supports `lightsail` as a provider")
	}

	return lightsail.NewDNSProvider()
}
