//go:build lego_katapult

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/katapult"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "katapult" {
		return nil, fmt.Errorf("this build of legotapas only supports `katapult` as a provider")
	}

	return katapult.NewDNSProvider()
}
