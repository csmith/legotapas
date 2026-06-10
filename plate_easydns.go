//go:build lego_easydns

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/easydns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "easydns" {
		return nil, fmt.Errorf("this build of legotapas only supports `easydns` as a provider")
	}

	return easydns.NewDNSProvider()
}
