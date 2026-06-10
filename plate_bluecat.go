//go:build lego_bluecat

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/bluecat"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "bluecat" {
		return nil, fmt.Errorf("this build of legotapas only supports `bluecat` as a provider")
	}

	return bluecat.NewDNSProvider()
}
