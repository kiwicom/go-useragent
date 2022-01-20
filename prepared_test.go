package useragent

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fromModule(t *testing.T) {
	mod := debug.Module{
		Path:    "foo/bar/baz",
		Version: "v1.2.3-test",
		Sum:     "doesn't matter",
		Replace: nil,
	}

	want := UserAgent{
		Name:        "baz",
		Version:     "v1.2.3-test",
		Environment: "dev", // hardcoded
		SystemInfo:  fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
	}
	got := fromModule(mod)
	assert.NoError(t, got.Validate(), "generated UserAgent should be valid")
	assert.Equal(t, want, got, "should generate expected UserAgent")

}

func TestTesting(t *testing.T) {
	want := UserAgent{
		Name:        "TestTesting", // name of the current test
		Version:     runtime.Version(),
		Environment: "testing",
		SystemInfo:  goSystemInfo(),
	}
	got := Testing(t)
	assert.NoError(t, got.Validate(), "generated UserAgent should be valid")
	assert.Equal(t, want, got, "should generate expected UserAgent")
}
