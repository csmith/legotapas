//go:build lego_scannet

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/scannet"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "scannet" {
		return nil, fmt.Errorf("this build of legotapas only supports `scannet` as a provider")
	}

	return scannet.NewDNSProvider()
}
