//go:build lego_tencentcloud

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/tencentcloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "tencentcloud" {
		return nil, fmt.Errorf("this build of legotapas only supports `tencentcloud` as a provider")
	}

	return tencentcloud.NewDNSProvider()
}
