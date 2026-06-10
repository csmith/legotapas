//go:build lego_nederhost

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/nederhost"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "nederhost" {
		return nil, fmt.Errorf("this build of legotapas only supports `nederhost` as a provider")
	}

	return nederhost.NewDNSProvider()
}
