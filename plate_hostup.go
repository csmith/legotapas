//go:build lego_hostup

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/hostup"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "hostup" {
		return nil, fmt.Errorf("this build of legotapas only supports `hostup` as a provider")
	}

	return hostup.NewDNSProvider()
}
