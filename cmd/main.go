package main

import (
	"fmt"
	sf "github.com/mikeconroy/gosf/salesforce"
	"github.com/spf13/viper"
)

/*
 * gosf query "SELECT id FROM Account"
 *
 */
func main() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config.")
		fmt.Println(err)
		return
	}

	sfInstanceUrl := viper.GetString("SF_INSTANCE")
	fmt.Println(sfInstanceUrl)
	sfInstance := sf.Salesforce{
		InstanceUrl: sfInstanceUrl,
	}
	versionResp, _ := sfInstance.GetVersions()
	for _, version := range versionResp {
		fmt.Printf("Label: %s\tURL: %s\tVersion: %s\n", version.Label, version.Url, version.Version)
	}

	latestVersion, _ := sfInstance.GetLatestVersion()
	fmt.Println("Latest Version: " + latestVersion.Version)

	sfInstance2 := sf.New(sfInstanceUrl,
		"username", "password", "securityToken", "consumerKey", "consumerSecret")
	fmt.Println(sfInstance2)
}
