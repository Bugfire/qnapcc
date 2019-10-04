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

func init() {
	rootCmd.AddCommand(cmdLogout())
}

func cmdLogout() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "logout",
		Short: "Logout",
		Long:  "Logout QNAP NAS",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := logout()
			return err
		},
	}
	return cmd
}

func logout() error {
	qnapUrl := viper.GetString("Url")
	cookiesJson := viper.GetString("Cookies")
	var cookies []*http.Cookie
	err := json.Unmarshal([]byte(cookiesJson), &cookies)
	if err != nil {
		log.Fatal(err)
	}
	logoutResult, err := api.Logout(qnapUrl, cookies)
	if err != nil {
		log.Fatal(err)
	}

	if logoutResult.Error != nil {
		if logoutResult.Error.Code == 401 {
			// Unauthorized
		} else {
			fmt.Printf("Code=%d ", logoutResult.Error.Code)
			fmt.Println(logoutResult.Error.Message)
			os.Exit(1)
		}
	}

	viper.Set("Url", nil)
	viper.Set("Cookies", nil)

	return err
}
