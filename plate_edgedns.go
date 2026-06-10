//go:build lego_edgedns

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/edgedns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "edgedns" {
		return nil, fmt.Errorf("this build of legotapas only supports `edgedns` as a provider")
	}

	return edgedns.NewDNSProvider()
}
