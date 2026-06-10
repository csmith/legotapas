//go:build lego_vegadns

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/vegadns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "vegadns" {
		return nil, fmt.Errorf("this build of legotapas only supports `vegadns` as a provider")
	}

	return vegadns.NewDNSProvider()
}
