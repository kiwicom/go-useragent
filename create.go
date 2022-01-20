package useragent

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var createInputRegex = regexp.MustCompile(`^\S+$`)

// ErrNameInvalid is returned if the appName contains any whitespace characters
var ErrNameInvalid = errors.New("invalid appName")

// ErrVersionInvalid is returned if the version contains any whitespace characters
var ErrVersionInvalid = errors.New("invalid version")

// ErrEnvironmentInvalid is returned if the environment contains any whitespace characters
var ErrEnvironmentInvalid = errors.New("invalid environment")

// Format will generate a useragent string which is compliant with RFC 22.
func (ua *UserAgent) Format() (string, error) {
	err := ua.Validate()
	if err != nil {
		return "", err
	}

	var useragent strings.Builder

	compulsory := fmt.Sprintf("%s/%s (Kiwi.com %s)", ua.Name, ua.Version, ua.Environment)
	useragent.WriteString(compulsory)
	if ua.SystemInfo != "" {
		useragent.Grow(1 + len(ua.SystemInfo))
		useragent.WriteString(" ")
		useragent.WriteString(ua.SystemInfo)
	}

	return useragent.String(), nil
}

var _ fmt.Stringer = UserAgent{}

// String implements fmt.Stringer interface. It's a wrapper around
// UserAgent.Format method. The `<invalid>` string is returned if any error
// occurs during generation.
func (ua UserAgent) String() string {
	s, err := ua.Format()
	if err != nil {
		return "<invalid>"
	}
	return s
}

// MustFormat is like UserAgent.Format but panics if any errors occurred during generation.
// It simplifies safe initialization of global variables holding default or prepared values of UserAgent.
func (ua UserAgent) MustFormat() string {
	s, err := ua.Format()
	if err != nil {
		panic(err)
	}
	return s
}
