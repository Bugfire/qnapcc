package api

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

var (
	ApiV1 string = "/containerstation/api/v1/"

	LoginPath  = "login"
	LogoutPath = "logout"

	SystemPath                  = "system"
	SystemResourcePath          = "system/resource"
	SystemReportPath            = "system/report"
	SystemBridgePath            = "system/bridge"
	TlsPath                     = "tls"
	TlsExportPath               = "tls/export"
	TlsDomainNamesPath          = "tls/domain_names"
	TlsExportRegistryPath       = "tls/export/registry"
	SharefolderPath             = "sharefolder/"
	ContainerPath               = "container"
	ImportConfigPath            = "import/config"
	BackgroundPath              = "background/"
	AppsPath                    = "apps"
	ResourceDevicePath          = "resource/device"
	ImagePath                   = "image"
	RegistryPath                = "registry"
	RegistryPushPath            = "registry/push"
	RegistryPingPath            = "registry/ping"
	NetworksPath                = "networks"
	VolumesPath                 = "volumes"
	LogPath                     = "log"
	LogExportPath               = "log/export"
	EventPath                   = "event"
	PreferencePath              = "preference/"
	WizardWorkspacePath         = "wizard/workspace"
	WizardWorkspaceStatusPath   = "wizard/workspace/status"
	PreferenceRepoPath          = "preference/repo"
	PreferenceNetworkPath       = "preference/network"
	PreferenceNetworkDockerPath = "preference/network_docker"
)

type ErrorDef struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseResult struct {
	Error *ErrorDef `json:"error"` // nil on success
}

func dumpValue(value interface{}, prefix string, key string) error {
	header := ""
	if prefix != "" || key != "" {
		header = fmt.Sprintf("%-24s", prefix+key+":")
	}
	switch value.(type) {
	case map[string]interface{}:
		if header != "" {
			fmt.Printf("%s\n", header)
			prefix = "  " + prefix
		}
		dumpDict(value.(map[string]interface{}), prefix)
		break
	case []interface{}:
		if header != "" {
			fmt.Printf("%s\n", header)
			prefix = "  " + prefix
		}
		dumpArray(value.([]interface{}), prefix)
		break
	case string:
		fmt.Printf("%s%s\n", header, strings.Replace(value.(string), "\n", "\\n", -1))
		break
	case float64:
		fmt.Printf("%s%f\n", header, value.(float64))
		break
	case int:
		fmt.Printf("%s%d\n", header, value.(int))
		break
	case bool:
		fmt.Printf("%s%t\n", header, value.(bool))
		break
	case nil:
		fmt.Printf("%snull\n", header)
		break
	default:
		fmt.Printf("%s?? %s\n", header, reflect.TypeOf(value))
		break
	}
	return nil
}

func dumpArray(array []interface{}, prefix string) error {
	for key, value := range array {
		str := fmt.Sprintf("[%d]", key)
		dumpValue(value, prefix, str)
	}
	return nil
}

func dumpDict(dict map[string]interface{}, prefix string) error {
	keys := []string{}
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		dumpValue(dict[key], prefix, key)
	}
	return nil
}

func Dump(value interface{}) error {
	if err := dumpValue(value, "", ""); err != nil {
		return err
	}
	return nil
}
