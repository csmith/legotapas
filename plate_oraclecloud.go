//go:build lego_oraclecloud

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/oraclecloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "oraclecloud" {
		return nil, fmt.Errorf("this build of legotapas only supports `oraclecloud` as a provider")
	}

	return oraclecloud.NewDNSProvider()
}
