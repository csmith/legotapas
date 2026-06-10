//go:build lego_zonomi

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/zonomi"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "zonomi" {
		return nil, fmt.Errorf("this build of legotapas only supports `zonomi` as a provider")
	}

	return zonomi.NewDNSProvider()
}
