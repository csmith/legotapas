//go:build lego_wannafind

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/wannafind"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "wannafind" {
		return nil, fmt.Errorf("this build of legotapas only supports `wannafind` as a provider")
	}

	return wannafind.NewDNSProvider()
}
