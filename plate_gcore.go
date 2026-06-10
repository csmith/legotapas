//go:build lego_gcore

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/gcore"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "gcore" {
		return nil, fmt.Errorf("this build of legotapas only supports `gcore` as a provider")
	}

	return gcore.NewDNSProvider()
}
