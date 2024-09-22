package salesforce

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (sf Salesforce) Authenticate() (string, error) {
	client := &http.Client{}
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", sf.ConsumerKey)
	data.Set("client_secret", sf.ConsumerSecret)
	data.Set("username", sf.Username)
	fmt.Println(sf.Password + sf.SecurityToken)
	data.Set("password", sf.Password+sf.SecurityToken)
	req, _ := http.NewRequest("POST", sf.InstanceUrl+"/services/oauth2/token", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	fmt.Println(res)
	b, _ := io.ReadAll(res.Body)
	fmt.Println(string(b))
	return "", nil
}
