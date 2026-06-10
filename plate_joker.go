//go:build lego_joker

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/joker"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "joker" {
		return nil, fmt.Errorf("this build of legotapas only supports `joker` as a provider")
	}

	return joker.NewDNSProvider()
}
