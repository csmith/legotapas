//go:build lego_acmedns

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/acmedns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "acmedns" {
		return nil, fmt.Errorf("this build of legotapas only supports `acmedns` as a provider")
	}

	return acmedns.NewDNSProvider()
}
