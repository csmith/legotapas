//go:build lego_curanet

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/curanet"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "curanet" {
		return nil, fmt.Errorf("this build of legotapas only supports `curanet` as a provider")
	}

	return curanet.NewDNSProvider()
}
