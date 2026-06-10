//go:build lego_freemyip

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/freemyip"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "freemyip" {
		return nil, fmt.Errorf("this build of legotapas only supports `freemyip` as a provider")
	}

	return freemyip.NewDNSProvider()
}
