//go:build lego_dreamhost

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/dreamhost"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "dreamhost" {
		return nil, fmt.Errorf("this build of legotapas only supports `dreamhost` as a provider")
	}

	return dreamhost.NewDNSProvider()
}
