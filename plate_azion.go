//go:build lego_azion

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/azion"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "azion" {
		return nil, fmt.Errorf("this build of legotapas only supports `azion` as a provider")
	}

	return azion.NewDNSProvider()
}
