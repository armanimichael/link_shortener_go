package http_utils

import (
	"net/url"
)

// IsValidUrl validates an url based on the structure of the passed url
// it does not consider the existence of an actual server
func IsValidUrl(rawUrl string) bool {
	parsedUrl, e := url.Parse(rawUrl)
	return e == nil && parsedUrl.Scheme != "" && parsedUrl.Host != ""
}
