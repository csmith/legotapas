//go:build lego_zilore

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/zilore"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "zilore" {
		return nil, fmt.Errorf("this build of legotapas only supports `zilore` as a provider")
	}

	return zilore.NewDNSProvider()
}
