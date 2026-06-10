//go:build lego_simply

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/simply"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "simply" {
		return nil, fmt.Errorf("this build of legotapas only supports `simply` as a provider")
	}

	return simply.NewDNSProvider()
}
