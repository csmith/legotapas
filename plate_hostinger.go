//go:build lego_hostinger

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/hostinger"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "hostinger" {
		return nil, fmt.Errorf("this build of legotapas only supports `hostinger` as a provider")
	}

	return hostinger.NewDNSProvider()
}
