package useragent

import "fmt"

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

// Validate an ua struct values and ensure that Format call generate a valid
// useragent. Descriptive non-nil error will be returned if any field of the ua
// contains value that prevents generation of correct RFC 22 useragent.
func (ua UserAgent) Validate() error {
	appNameMatch := createInputRegex.MatchString(ua.Name)
	if !appNameMatch {
		return fmt.Errorf("error validating %q: %w", ua.Name, ErrNameInvalid)
	}
	versionMatch := createInputRegex.MatchString(ua.Version)
	if !versionMatch {
		return fmt.Errorf("error validating %q: %w", ua.Version, ErrVersionInvalid)
	}
	environmentMatch := createInputRegex.MatchString(ua.Environment)
	if !environmentMatch {
		return fmt.Errorf("error validating %q: %w", ua.Environment, ErrEnvironmentInvalid)
	}
	return nil
}
