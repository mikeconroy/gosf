package salesforce

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	tokenPath = "/services/oauth2/token"
)

type AuthenticateResponse struct {
	AccessToken string `json:"access_token"`
	InstanceUrl string `json:"instance_url"`
	Id          string `json:"id"`
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`
}

type Authenticator interface {
	Authenticate(Salesforce) (AuthenticateResponse, error)
}

type UsernamePasswordAuthenticator struct {
	Username       string
	Password       string
	SecurityToken  string
	ConsumerKey    string
	ConsumerSecret string
}

func (auth UsernamePasswordAuthenticator) Authenticate(sf Salesforce) (AuthenticateResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", auth.ConsumerKey)
	data.Set("client_secret", auth.ConsumerSecret)
	data.Set("username", auth.Username)
	data.Set("password", auth.Password+auth.SecurityToken)
	req, err := http.NewRequest("POST", sf.InstanceUrl+tokenPath, strings.NewReader(data.Encode()))
	if err != nil {
		return AuthenticateResponse{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := sf.HttpClient.Do(req)
	if err != nil {
		return AuthenticateResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var result AuthenticateResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cannot unmarshal JSON")
	}
	return result, err
}

type ClientCredentialsAuthenticator struct {
	ConsumerKey    string
	ConsumerSecret string
}

func (auth ClientCredentialsAuthenticator) Authenticate(sf Salesforce) (AuthenticateResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", auth.ConsumerKey)
	data.Set("client_secret", auth.ConsumerSecret)
	req, err := http.NewRequest("POST", sf.InstanceUrl+tokenPath, strings.NewReader(data.Encode()))
	if err != nil {
		return AuthenticateResponse{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := sf.HttpClient.Do(req)
	if err != nil {
		return AuthenticateResponse{}, err
	}
	defer resp.Body.Close()
	// TODO: Handle errors - check response status code & payload
	// {"error":"invalid_grant","error_description":"no client credentials user enabled"}

	body, err := io.ReadAll(resp.Body)

	var result AuthenticateResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cannot unmarshal JSON")
	}
	return result, err
}
