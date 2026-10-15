//go:build lego_feno

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/feno"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "feno" {
		return nil, fmt.Errorf("this build of legotapas only supports `feno` as a provider")
	}

	return feno.NewDNSProvider()
}
