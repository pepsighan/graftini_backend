package sanitize

import "strings"

// CleanRoute removes any trailing slashes from the route.
func CleanRoute(path string) string {
	return strings.TrimSuffix(path, "/")
}
