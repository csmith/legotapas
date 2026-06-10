//go:build lego_dnsupdate

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/dnsupdate"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "dnsupdate" {
		return nil, fmt.Errorf("this build of legotapas only supports `dnsupdate` as a provider")
	}

	return dnsupdate.NewDNSProvider()
}
