//go:build lego_abion

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/abion"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "abion" {
		return nil, fmt.Errorf("this build of legotapas only supports `abion` as a provider")
	}

	return abion.NewDNSProvider()
}
