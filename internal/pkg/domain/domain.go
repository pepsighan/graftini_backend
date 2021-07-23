package domain

import (
	"fmt"

	"github.com/pepsighan/graftini_backend/internal/pkg/config"
)

// GenerateDomainNameFromRefID generates a full domain name from the Ref ID of the project.
func GenerateDomainNameFromRefID(refID string, env config.Environment) string {
	return fmt.Sprintf("%v.%v", refID, suffixDomainName(env))
}

// suffixDomainName gives the domain suffix to use. We use graftini.app for production
// and graftini.xyz for development.
func suffixDomainName(env config.Environment) string {
	if env.IsProduction() {
		return "graftini.app"
	}

	// Since the same domain is being used for development and local, the graftini.xyz is
	// divided by subdomain.
	return fmt.Sprintf("%v.graftini.xyz", env)
}
