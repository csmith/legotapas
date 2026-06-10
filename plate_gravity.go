//go:build lego_gravity

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/gravity"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "gravity" {
		return nil, fmt.Errorf("this build of legotapas only supports `gravity` as a provider")
	}

	return gravity.NewDNSProvider()
}
