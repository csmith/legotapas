//go:build lego_onlinenet

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/onlinenet"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "onlinenet" {
		return nil, fmt.Errorf("this build of legotapas only supports `onlinenet` as a provider")
	}

	return onlinenet.NewDNSProvider()
}
