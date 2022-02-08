//go:build !lego_internetbs && !lego_rfc2136 && !lego_bindman && !lego_cloudxns && !lego_alidns && !lego_glesys && !lego_infoblox && !lego_namesilo && !lego_gcore && !lego_httpreq && !lego_digitalocean && !lego_stackpath && !lego_epik && !lego_designate && !lego_vultr && !lego_gandiv5 && !lego_porkbun && !lego_simply && !lego_auroradns && !lego_oraclecloud && !lego_exoscale && !lego_duckdns && !lego_vinyldns && !lego_ibmcloud && !lego_ionos && !lego_zonomi && !lego_lightsail && !lego_njalla && !lego_iij && !lego_hyperone && !lego_edgedns && !lego_acmedns && !lego_desec && !lego_mydnsjp && !lego_wedos && !lego_hosttech && !lego_linode && !lego_cloudflare && !lego_mythicbeasts && !lego_ns1 && !lego_gcloud && !lego_loopia && !lego_bluecat && !lego_luadns && !lego_vegadns && !lego_scaleway && !lego_dyn && !lego_namedotcom && !lego_tencentcloud && !lego_vscale && !lego_cloudns && !lego_transip && !lego_infomaniak && !lego_dode && !lego_dnspod && !lego_hostingde && !lego_sonic && !lego_dnsimple && !lego_liquidweb && !lego_autodns && !lego_servercow && !lego_joker && !lego_route53 && !lego_azure && !lego_hurricane && !lego_inwx && !lego_versio && !lego_gandi && !lego_conoha && !lego_sakuracloud && !lego_dynu && !lego_pdns && !lego_rimuhosting && !lego_domeneshop && !lego_regru && !lego_selectel && !lego_zoneee && !lego_arvancloud && !lego_nicmanager && !lego_rackspace && !lego_freemyip && !lego_clouddns && !lego_otc && !lego_allinkl && !lego_nifcloud && !lego_easydns && !lego_dreamhost && !lego_exec && !lego_checkdomain && !lego_constellix && !lego_yandex && !lego_namecheap && !lego_dnsmadeeasy && !lego_netcup && !lego_netlify && !lego_hetzner && !lego_godaddy && !lego_safedns && !lego_ovh

package legotapas

import (
    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    return dns.NewDNSChallengeProviderByName(providerName)
}
