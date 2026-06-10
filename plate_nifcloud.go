//go:build lego_nifcloud

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/nifcloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "nifcloud" {
		return nil, fmt.Errorf("this build of legotapas only supports `nifcloud` as a provider")
	}

	return nifcloud.NewDNSProvider()
}
