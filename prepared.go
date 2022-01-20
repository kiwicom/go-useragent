package useragent

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"runtime/debug"
	"testing"
)

// FromModule generated UserAgent based on debug.ReadBuildInfo.
// If no build info available (e.g. in tests) the error will be returned.
func FromModule() (UserAgent, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return UserAgent{}, errors.New("cannot retrieve build info")
	}
	return fromModule(info.Main), nil
}

func fromModule(mod debug.Module) UserAgent {
	_, moduleBase := path.Split(mod.Path) // get the last element of the full module path
	return UserAgent{
		Name:        moduleBase,
		Version:     mod.Version,
		Environment: "dev",
		SystemInfo:  goSystemInfo(),
	}
}

// Testing is a helper constructor to help with building UserAgent for tests.
func Testing(t testing.TB) UserAgent {
	t.Helper()
	return UserAgent{
		Name:        t.Name(),
		Version:     runtime.Version(),
		Environment: "testing",
		SystemInfo:  goSystemInfo(),
	}
}

func goSystemInfo() string {
	return fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
