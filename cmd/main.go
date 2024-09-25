package main

import (
	"fmt"
	sf "github.com/mikeconroy/gosf/salesforce"
	"github.com/spf13/viper"
)

type Config struct {
	SfInstanceUrl    string `mapstructure:"SF_INSTANCE"`
	SfUsername       string `mapstructure:"SF_USERNAME"`
	SfPassword       string `mapstructure:"SF_PASSWORD"`
	SfConsumerKey    string `mapstructure:"SF_CONSUMER_KEY"`
	SfConsumerSecret string `mapstructure:"SF_CONSUMER_SECRET"`
	SfSecurityToken  string `mapstructure:"SF_SECURITY_TOKEN"`
}

/*
 * gosf query "SELECT id FROM Account"
 *
 */
func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Failed to load config.", err)
	}

	sfInstance := sf.New(config.SfInstanceUrl,
		config.SfUsername,
		config.SfPassword,
		config.SfSecurityToken,
		config.SfConsumerKey,
		config.SfConsumerSecret)

	sfInstance.Authenticate()
	queryRes, err := sfInstance.Query("SELECT Id,Name FROM Account")
	for _, record := range queryRes.Records {
		fmt.Println("ID:", record["Id"], "Name:", record["Name"])
	}
}

func printVersions(sfInstance sf.Salesforce) {
	versionResp, _ := sfInstance.GetVersions()
	for _, version := range versionResp {
		fmt.Printf("Label: %s\tURL: %s\tVersion: %s\n", version.Label, version.Url, version.Version)
	}
	latestVersion, _ := sfInstance.GetLatestVersion()
	fmt.Println("Latest Version: " + latestVersion.Version)
}

func loadConfig() (Config, error) {
	config := Config{}
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	// sfInstanceUrl := viper.GetString("SF_INSTANCE")
	// fmt.Println(sfInstanceUrl)
	return config, nil
}
