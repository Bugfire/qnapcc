package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bugfire/qnapcc/api"
)

func init() {
	rootCmd.AddCommand(cmdLogin())
}

func cmdLogin() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "login <URL>",
		Short: "Login",
		Long:  "Login specified QNAP NAS",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := login(args[0])
			return err
		},
	}
	return cmd
}

func login(qnapUrl string) error {
	fmt.Print("User: ")
	var username string
	fmt.Scan(&username)
	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\n")

	loginResult, cookies, err := api.Login(qnapUrl, username, string(password))
	if err != nil {
		log.Fatal(err)
	}
	if loginResult.Error != nil {
		fmt.Printf("Code=%d ", loginResult.Error.Code)
		fmt.Println(loginResult.Error.Message)
		os.Exit(1)
	}

	cookiesJson, err := json.Marshal(cookies)
	if err != nil {
		log.Fatal(err)
	}
	viper.Set("Url", qnapUrl)
	viper.Set("Cookies", string(cookiesJson))
	if viper.ConfigFileUsed() != "" {
		if err := viper.WriteConfig(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := viper.WriteConfigAs("./.qnapcc.yml"); err != nil {
			log.Fatal(err)
		}
	}

	return err
}
