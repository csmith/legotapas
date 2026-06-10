//go:build lego_huaweicloud

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/huaweicloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "huaweicloud" {
		return nil, fmt.Errorf("this build of legotapas only supports `huaweicloud` as a provider")
	}

	return huaweicloud.NewDNSProvider()
}
