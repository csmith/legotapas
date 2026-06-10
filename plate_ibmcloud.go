//go:build lego_ibmcloud

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/ibmcloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "ibmcloud" {
		return nil, fmt.Errorf("this build of legotapas only supports `ibmcloud` as a provider")
	}

	return ibmcloud.NewDNSProvider()
}
