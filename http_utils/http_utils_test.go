package http_utils

import "testing"

func Test_IsValidUrl_Valid(t *testing.T) {
	url := "https://www.duckduckgo.com"
	if !IsValidUrl(url) {
		t.Errorf("this url should be valid")
	}
}

func Test_IsValidUrl_Invalid(t *testing.T) {
	url := "nonvalid"
	if IsValidUrl(url) {
		t.Errorf("this url should NOT be valid")
	}
}
