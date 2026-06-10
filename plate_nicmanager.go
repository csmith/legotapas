//go:build lego_nicmanager

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/nicmanager"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "nicmanager" {
		return nil, fmt.Errorf("this build of legotapas only supports `nicmanager` as a provider")
	}

	return nicmanager.NewDNSProvider()
}
