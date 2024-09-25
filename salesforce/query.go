package salesforce

import (
	"encoding/json"
	"io"
	"net/http"
)

type QueryResponse struct {
	Done      bool             `json:"done"`
	TotalSize int              `json:"totalSize"`
	Records   []map[string]any `json:"records"`
}

/*
 * https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/dome_query.htm
 */
func (sf Salesforce) Query(soqlQuery string) (QueryResponse, error) {
	req, err := http.NewRequest("GET", sf.InstanceUrl+sf.ApiVersionPath+"/query", nil)
	if err != nil {
		return QueryResponse{}, err
	}
	q := req.URL.Query()
	q.Add("q", soqlQuery)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", "Bearer "+sf.Token)
	resp, err := sf.HttpClient.Do(req)
	if err != nil {
		return QueryResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var result QueryResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return QueryResponse{}, err
	}
	return result, err

}
