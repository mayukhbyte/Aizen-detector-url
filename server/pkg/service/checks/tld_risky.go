package checks

import "github.com/urlvet/urlvet/pkg/constants"

func IsRiskyTld(domain string) (bool, bool, string) {
	tld, icann := GetTld(domain)
	_, ok := constants.RiskyTLDs[tld]

	return ok, icann, tld
}
