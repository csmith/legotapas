//go:build lego_nexdns

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/nexdns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "nexdns" {
		return nil, fmt.Errorf("this build of legotapas only supports `nexdns` as a provider")
	}

	return nexdns.NewDNSProvider()
}
