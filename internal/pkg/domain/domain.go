package domain

import (
	"fmt"

	"github.com/pepsighan/graftini_backend/internal/pkg/config"
)

// GenerateDomainNameFromRefID generates a full domain name from the Ref ID of the project.
func GenerateDomainNameFromRefID(refID string, env config.Environment) string {
	return fmt.Sprintf("%v.%v", refID, suffixDomainName(env))
}

const graftiniAppDomain string = "graftini.app"

// suffixDomainName gives the domain suffix to use. We use [graftiniaAppDomain]
// for production and for others we append with development & local.
func suffixDomainName(env config.Environment) string {
	if env.IsProduction() {
		// This is hard-coded for the app. There is no other domain we use.
		return graftiniAppDomain
	}

	return fmt.Sprintf("%v.%v", env, graftiniAppDomain)
}
