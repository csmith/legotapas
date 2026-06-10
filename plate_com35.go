//go:build lego_com35

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/com35"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "com35" {
		return nil, fmt.Errorf("this build of legotapas only supports `com35` as a provider")
	}

	return com35.NewDNSProvider()
}
