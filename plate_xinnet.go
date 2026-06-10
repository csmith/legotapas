//go:build lego_xinnet

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/xinnet"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "xinnet" {
		return nil, fmt.Errorf("this build of legotapas only supports `xinnet` as a provider")
	}

	return xinnet.NewDNSProvider()
}
