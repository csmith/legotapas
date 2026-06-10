//go:build lego_keyhelp

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/keyhelp"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "keyhelp" {
		return nil, fmt.Errorf("this build of legotapas only supports `keyhelp` as a provider")
	}

	return keyhelp.NewDNSProvider()
}
