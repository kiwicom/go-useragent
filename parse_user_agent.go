// Package useragent provides tools to work with user agent strings as used in Kiwi.com
package useragent

import (
	"errors"
	"regexp"
)

// UserAgent contains all available fields in UserAgent according to RFC-22 rules.
type UserAgent struct {
	// Name is the service name
	Name string
	// Version is the version of the remote service, can be tag or commit hash
	Version string
	// Environment decribes where the service is run (e.g. sandbox, production)
	Environment string
	// SystemInfo describes extra metadata (e.g. node-fetch 1.2)
	SystemInfo string
}

// ErrInvalidFormat is returned when the regex parse fails. This means the User Agent did not adhere to RFC-22 rules.
var ErrInvalidFormat = errors.New("user agent string is not compliant with Kiwi RFC-22")

var userAgentRegex = regexp.MustCompile(`^(?P<name>\S+?)\/(?P<version>\S+?) \(Kiwi\.com (?P<environment>\S+?)\)(?: ?(?P<system_info>.*))$`)

// Parse parses userAgent string according to RFC-22 rules.
func Parse(userAgent string) (UserAgent, error) {
	match := userAgentRegex.FindStringSubmatch(userAgent)
	if match == nil {
		return UserAgent{}, ErrInvalidFormat
	}
	ua := UserAgent{
		Name:        match[1],
		Version:     match[2],
		Environment: match[3],
		SystemInfo:  match[4],
	}
	return ua, nil
}
