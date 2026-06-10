//go:build lego_onecloudru

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/onecloudru"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "onecloudru" {
		return nil, fmt.Errorf("this build of legotapas only supports `onecloudru` as a provider")
	}

	return onecloudru.NewDNSProvider()
}
