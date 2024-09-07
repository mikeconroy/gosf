package main

import (
	"fmt"
	sf "github.com/mikeconroy/gosf/salesforce"
)

/*
 * gosf query "SELECT id FROM Account"
 *
 */
func main() {
	sf := sf.Salesforce{
		InstanceUrl: "https://{INSTANCEID}.develop.my.salesforce.com",
	}
	versionResp, _ := sf.GetVersions()
	for _, version := range versionResp {
		fmt.Printf("Label: %s\tURL: %s\tVersion: %s\n", version.Label, version.Url, version.Version)
	}

	latestVersion, _ := sf.GetLatestVersion()
	fmt.Println("Latest Version: " + latestVersion.Version)

}
