package build

import (
	"fmt"
	"io/ioutil"
)

func WriteConfig() error {
	version, err := GetVersion()
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("server/buildconfig/config.go", []byte(fmt.Sprintf(`
package buildconfig

var (
	Config BuildConfig = BuildConfig {
		Product: "Mars Decimation",
		Version: "%s",
	}
)
`, version)), 0644); err != nil {
		return err
	}
	return nil
}
