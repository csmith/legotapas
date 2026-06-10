//go:build lego_fornex

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/fornex"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "fornex" {
		return nil, fmt.Errorf("this build of legotapas only supports `fornex` as a provider")
	}

	return fornex.NewDNSProvider()
}
