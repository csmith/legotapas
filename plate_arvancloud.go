//go:build lego_arvancloud

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/arvancloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "arvancloud" {
		return nil, fmt.Errorf("this build of legotapas only supports `arvancloud` as a provider")
	}

	return arvancloud.NewDNSProvider()
}
