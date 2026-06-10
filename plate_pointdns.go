//go:build lego_pointdns

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/pointdns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "pointdns" {
		return nil, fmt.Errorf("this build of legotapas only supports `pointdns` as a provider")
	}

	return pointdns.NewDNSProvider()
}
