//go:build lego_vultr

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/vultr"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "vultr" {
		return nil, fmt.Errorf("this build of legotapas only supports `vultr` as a provider")
	}

	return vultr.NewDNSProvider()
}
