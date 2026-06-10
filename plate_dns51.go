//go:build lego_dns51

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/dns51"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "dns51" {
		return nil, fmt.Errorf("this build of legotapas only supports `dns51` as a provider")
	}

	return dns51.NewDNSProvider()
}
