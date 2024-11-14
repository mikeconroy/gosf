package salesforce

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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

type JWTBearerAuthenticator struct {
	ConsumerKey string
	Username    string
	PrivateKey  *rsa.PrivateKey
}

func (auth JWTBearerAuthenticator) Authenticate(sf Salesforce) (AuthenticateResponse, error) {
	jwt, _ := auth.GenerateJWT(sf)
	fmt.Println(jwt)
	return AuthenticateResponse{}, nil
}

func (auth JWTBearerAuthenticator) GenerateJWT(sf Salesforce) (string, error) {

	header := []byte("{\"alg\":\"RS256\"}")
	encodedHeader := base64.StdEncoding.EncodeToString(header)

	issuer := auth.ConsumerKey
	audience := sf.InstanceUrl
	subject := auth.Username
	expiry := fmt.Sprint(time.Now().Add(15 * time.Minute).Unix())
	payload := []byte("{" +
		"\"iss\":\"" + issuer + "\"," +
		"\"aud\":\"" + audience + "\"," +
		"\"sub\":\"" + subject + "\"," +
		"\"exp\":\"" + expiry + "\"}")
	encodedPayload := base64.StdEncoding.EncodeToString(payload)

	hash := sha256.New()
	hash.Write(payload)
	digest := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, auth.PrivateKey, crypto.SHA256, digest)
	if err != nil {
		return "", err
	}
	encodedSignature := base64.StdEncoding.EncodeToString(signature)

	token := encodedHeader + "." + encodedPayload + "." + encodedSignature
	return token, err
}
