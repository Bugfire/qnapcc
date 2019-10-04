package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bugfire/qnapcc/api"
)

type dumpDef struct {
	Path   string
	Use    string
	Short  string
	Long   string
	IsJson bool
}

func init() {
	dumps := []dumpDef{
		{
			Path:   api.SystemPath,
			Use:    "describe-system",
			Short:  "Describe system",
			Long:   "Describe system information",
			IsJson: true,
		},
		{
			Path:   api.SystemResourcePath,
			Use:    "describe-system-resource",
			Short:  "Describe system resource",
			Long:   "Describe system resource information",
			IsJson: true,
		},
		{
			Path:   api.SystemReportPath,
			Use:    "download-system-report",
			Short:  "Download system report as tar.bz2",
			Long:   "Download system diagnosis report as tar.bz2",
			IsJson: false,
		},
		{
			Path:   api.SystemBridgePath,
			Use:    "describe-system-bridge",
			Short:  "Describe bridge",
			Long:   "Describe brief bridge information",
			IsJson: true,
		},
		{
			Path:   api.TlsPath,
			Use:    "describe-tls",
			Short:  "Describe TLS",
			Long:   "Describe certificate information",
			IsJson: true,
		},
		{
			Path:   api.TlsDomainNamesPath,
			Use:    "describe-domain-names",
			Short:  "Describe domain names",
			Long:   "Describe extra DNS hostname or IP Address for server certificate",
			IsJson: true,
		},
		{
			Path:   api.TlsExportPath,
			Use:    "download-tls",
			Short:  "Download certificate files as ZIP",
			Long:   "Download certificate files as ZIP format",
			IsJson: false,
		},
		{
			Path:   api.TlsExportRegistryPath,
			Use:    "download-tls-registry",
			Short:  "Download certificate file for registry",
			Long:   "Download certificate file for registry",
			IsJson: false,
		},
		{
			Path:   api.SharefolderPath,
			Use:    "list-shared-folders",
			Short:  "List shared folders",
			Long:   "List shared folders",
			IsJson: true,
		},
		{
			Path:   api.ContainerPath,
			Use:    "list-containers",
			Short:  "List containers",
			Long:   "List containers",
			IsJson: true,
		},
		{
			Path:   api.ImportConfigPath,
			Use:    "describe-import-config",
			Short:  "Describe given container archive path, query the configure",
			Long:   "Describe given container archive path, query the configure",
			IsJson: true,
		},
		{
			Path:   api.BackgroundPath,
			Use:    "list-tasks",
			Short:  "List background tasks",
			Long:   "List background tasks",
			IsJson: true,
		},
		{
			Path:   api.AppsPath,
			Use:    "list-apps",
			Short:  "List all custom application information",
			Long:   "List all custom application information",
			IsJson: true,
		},
		{
			Path:   api.ResourceDevicePath,
			Use:    "list-devices",
			Short:  "List available device list. The device allows access inside container",
			Long:   "List available device list. The device allows access inside container",
			IsJson: true,
		},
		{
			Path:   api.ImagePath,
			Use:    "list-images",
			Short:  "List recommended image list. If depots is from dockerhub, it will take a few seconds to search.",
			Long:   "List recommended image list. If depots is from dockerhub, it will take a few seconds to search.",
			IsJson: true,
		},
		{
			Path:   api.RegistryPath,
			Use:    "list-registry",
			Short:  "List registry",
			Long:   "List registry",
			IsJson: true,
		},
		{
			Path:   api.RegistryPushPath,
			Use:    "describe-default-push-registry",
			Short:  "Describe default push registry",
			Long:   "Describe default push registry",
			IsJson: true,
		},
		/*
			{
				Path:   api.RegistryPingPath,
				Use:    "ping-registry",
				Short:  "Ping a registry",
				Long:   "Ping a registry",
				IsJson: false,
			},
		*/
		{
			Path:   api.NetworksPath,
			Use:    "list-networks",
			Short:  "List all networks",
			Long:   "List all networks",
			IsJson: true,
		},
		{
			Path:   api.VolumesPath,
			Use:    "list-volumes",
			Short:  "List all volumes",
			Long:   "List all volumes",
			IsJson: true,
		},
		{
			Path:   api.LogPath,
			Use:    "list-log",
			Short:  "List log and event",
			Long:   "List log and event",
			IsJson: true,
		},
		{
			Path:   api.LogExportPath,
			Use:    "download-log",
			Short:  "Download log as CSV",
			Long:   "Download log as CSV",
			IsJson: false,
		},
		{
			Path:   api.EventPath,
			Use:    "list-event",
			Short:  "List event",
			Long:   "List event",
			IsJson: true,
		},
		{
			Path:   api.PreferencePath,
			Use:    "list-prefs",
			Short:  "List preference",
			Long:   "List preference",
			IsJson: true,
		},
		{
			Path:   api.WizardWorkspacePath,
			Use:    "describe-wizard-workspace",
			Short:  "Describe wizard workspace",
			Long:   "Describe wizard workspace",
			IsJson: true,
		},
		{
			Path:   api.WizardWorkspaceStatusPath,
			Use:    "describe-wizard-workspace-status",
			Short:  "Describe wizard workspace status",
			Long:   "Describe wizard workspace status",
			IsJson: true,
		},
		{
			Path:   api.PreferenceRepoPath,
			Use:    "list-image-repository",
			Short:  "List image repository",
			Long:   "List image repository",
			IsJson: true,
		},
		{
			Path:   api.PreferenceNetworkPath,
			Use:    "describe-network-settings",
			Short:  "Describe network settings",
			Long:   "Describe network settings",
			IsJson: true,
		},
		{
			Path:   api.PreferenceNetworkDockerPath,
			Use:    "describe-docker-network-settings",
			Short:  "Describe docker's network settings",
			Long:   "Describe docker's network settings",
			IsJson: true,
		},
	}

	for _, value := range dumps {
		path := value.Path
		isJson := value.IsJson
		var cmd = &cobra.Command{
			Use:   value.Use,
			Short: value.Short,
			Long:  value.Long,
			Args:  cobra.ExactArgs(0),
			RunE: func(cmd *cobra.Command, args []string) error {
				err := dump(path, isJson)
				return err
			},
		}
		rootCmd.AddCommand(cmd)
	}
}

func dump(path string, isJson bool) error {
	qnapUrl := viper.GetString("Url")
	cookiesJson := viper.GetString("Cookies")
	var cookies []*http.Cookie
	err := json.Unmarshal([]byte(cookiesJson), &cookies)
	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("%s%s%s", qnapUrl, api.ApiV1, path)
	body, err := api.Get(url, cookies)
	if err != nil {
		log.Fatal(err)
	}

	if isJson {
		var baseResult api.BaseResult
		if err := json.Unmarshal(body, &baseResult); err == nil {
			if baseResult.Error != nil {
				fmt.Printf("Code=%d ", baseResult.Error.Code)
				fmt.Println(baseResult.Error.Message)
				os.Exit(1)
			}
		}

		var anon interface{}
		if err := json.Unmarshal(body, &anon); err != nil {
			return err
		}
		api.Dump(anon)
	} else {
		fmt.Println(string(body))
	}

	return err
}
