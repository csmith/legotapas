//go:build lego_mijnhost

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/mijnhost"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "mijnhost" {
		return nil, fmt.Errorf("this build of legotapas only supports `mijnhost` as a provider")
	}

	return mijnhost.NewDNSProvider()
}
