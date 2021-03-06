// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package config

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

// DefaultConfigFileContent is the content of default config file.
const DefaultConfigFileContent = `# OIS services configuration

#access_key_id: ACCESS_KEY_ID
#secret_access_key: SECRET_ACCESS_KEY

host: ois.com
port: 443
protocol: https

# Additional User-Agent
additional_user_agent: ""

# Valid log levels are "debug", "info", "warn", "error", and "fatal".
log_level: warn

`

const (
	// DefaultConfigFile is the filename of default config file.
	DefaultConfigFile = "~/.ois/config.yaml"

	// EnvConfigPath is config environment variable.
	EnvConfigPath = "QINGSTOR_CONFIG_PATH"

	// EnvAccessKeyID is config envrionment variable.
	EnvAccessKeyID = "QINGSTOR_ACCESS_KEY_ID"

	// EnvSecretAccessKey is config envrionment variable.
	EnvSecretAccessKey = "QINGSTOR_SECRET_ACCESS_KEY"

	// EnvEnableVirtualHostStyle is config envrionment variable.
	EnvEnableVirtualHostStyle = "QINGSTOR_ENABLE_VIRTUAL_HOST_STYLE"

	// EnvEnableDualStack is config envrionment variable.
	EnvEnableDualStack = "QINGSTOR_ENABLE_DUAL_STACK"
)

// GetUserConfigFilePath returns the user config file path.
func GetUserConfigFilePath() string {
	configFile := DefaultConfigFile
	configPath := os.Getenv(EnvConfigPath)
	if configPath != "" {
		configFile = configPath
	}
	return strings.Replace(configFile, "~/", getHome()+"/", 1)
}

// InstallDefaultUserConfig will install default config file.
func InstallDefaultUserConfig() error {
	err := os.MkdirAll(path.Dir(GetUserConfigFilePath()), 0755)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(GetUserConfigFilePath(), []byte(DefaultConfigFileContent), 0644)
}

func getHome() string {
	home := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	}

	return home
}
