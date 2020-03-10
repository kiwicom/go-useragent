package useragent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	happyPathTests := map[UserAgent]string{
		UserAgent{Name: "test", Version: "1.0", Environment: "dev"}:                             "test/1.0 (Kiwi.com dev)",
		UserAgent{Name: "test-with-extra-stuff", Version: "1.0", Environment: "dev"}:            "test-with-extra-stuff/1.0 (Kiwi.com dev)",
		UserAgent{Name: "test", Version: "greqgr73423224", Environment: "dev"}:                  "test/greqgr73423224 (Kiwi.com dev)",
		UserAgent{Name: "test", Version: "1.0", Environment: "dev", SystemInfo: "requests 0.1"}: "test/1.0 (Kiwi.com dev) requests 0.1",
	}
	for test, expected := range happyPathTests {
		res, err := test.Format()
		assert.Equal(t, expected, res)
		assert.Equal(t, nil, err)
	}

	unhappyPathTests := map[UserAgent]error{
		UserAgent{Name: "test ", Version: "1.0", Environment: "dev"}:      ErrNameInvalid,
		UserAgent{Name: "test", Version: "1.0  ", Environment: "dev"}:     ErrVersionInvalid,
		UserAgent{Name: "test", Version: "1.0", Environment: "dev local"}: ErrEnvironmentInvalid,
	}
	for test, expected := range unhappyPathTests {
		res, err := test.Format()
		assert.Equal(t, expected, err)
		assert.Equal(t, "", res)
	}
}
