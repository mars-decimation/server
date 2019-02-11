package build

import (
	"fmt"
	"io/ioutil"
)

// WriteConfig writes the configuration file (server/buildconfig/config.go)
func WriteConfig() error {
	version, err := GetVersion()
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("server/buildconfig/config.go", []byte(fmt.Sprintf(`
package buildconfig

var (
	// Config is the configuration that is available at compile time
	Config = BuildConfig {
		Product: "Mars Decimation",
		Version: "%s",
	}
)
`, version)), 0644); err != nil {
		return err
	}
	return nil
}
