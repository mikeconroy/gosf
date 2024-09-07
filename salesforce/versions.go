package salesforce

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type VersionResponse struct {
	Label   string `json:"label"`
	Url     string `json:"url"`
	Version string `json:"version"`
}

func (sf Salesforce) GetVersions() ([]VersionResponse, error) {
	resp, err := http.Get(sf.InstanceUrl + "/services/data/")
	if resp == nil {
		fmt.Println("Get Versions API Call return nil.")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var result []VersionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cannot unmarshal JSON")
	}
	return result, err
}

func (sf Salesforce) GetLatestVersion() (VersionResponse, error) {
	versions, err := sf.GetVersions()
	return versions[len(versions)-1], err
}
