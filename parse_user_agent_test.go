package useragent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	happyPathTests := map[string]UserAgent{
		"BALKAN/4704b82 (Kiwi.com sandbox)":                                                                    UserAgent{Name: "BALKAN", Version: "4704b82", Environment: "sandbox", SystemInfo: ""},
		"balkan/1.42.1 (Kiwi.com sandbox)":                                                                     UserAgent{Name: "balkan", Version: "1.42.1", Environment: "sandbox", SystemInfo: ""},
		"TRANSACTIONAL_MESSAGING_-_NEST_APP/1.42.1 (Kiwi.com sandbox)":                                         UserAgent{Name: "TRANSACTIONAL_MESSAGING_-_NEST_APP", Version: "1.42.1", Environment: "sandbox", SystemInfo: ""},
		"flights/b5bf54c32a112273245543a9eac0afdb8d8b32a5 (Kiwi.com jimbobjim-dev) requests/2.16.3 python/2.7.15": UserAgent{Name: "flights", Version: "b5bf54c32a112273245543a9eac0afdb8d8b32a5", Environment: "jimbobjim-dev", SystemInfo: "requests/2.16.3 python/2.7.15"},
	}

	for test, expected := range happyPathTests {
		res, err := Parse(test)
		assert.Equal(t, expected, res)
		assert.Equal(t, nil, err)
	}

	unhappyPathTests := map[string]error{
		"":                                   ErrInvalidFormat,
		"balkan /1.42.1 (Kiwi.com sandbox)":  ErrInvalidFormat,
		" balkan /1.42.1 (Kiwi.com sandbox)": ErrInvalidFormat,
		"balkam":                             ErrInvalidFormat,
		"/4704b82 (Kiwi.com sandbox)":        ErrInvalidFormat,
		"wegf42314215csaz#@#$/!@#$":          ErrInvalidFormat,
		"  /5432 (Kiwi.com sandbox)":         ErrInvalidFormat,
		"stuff /  (Kiwi.com sandbox)":        ErrInvalidFormat,
		"balkan/1.42.1 ( sandbox)":           ErrInvalidFormat,
		"TRANSACTIONAL_MESSAGING_-_NEST_APP": ErrInvalidFormat,
	}

	for test, expected := range unhappyPathTests {
		res, err := Parse(test)
		assert.Equal(t, expected, err)
		assert.Equal(t, UserAgent{}, res)
	}
}
