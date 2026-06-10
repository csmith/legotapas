//go:build lego_webnamesru

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/webnamesru"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "webnamesru" {
		return nil, fmt.Errorf("this build of legotapas only supports `webnamesru` as a provider")
	}

	return webnamesru.NewDNSProvider()
}
