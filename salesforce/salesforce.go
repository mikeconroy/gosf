package salesforce

import (
	"net/http"
	"strconv"
	"time"
)

type Salesforce struct {
	InstanceUrl     string
	ApiVersionPath  string
	HttpClient      http.Client
	Authentication  Authenticator
	TokenValidFor   time.Duration
	Token           string
	TokenExpiryTime time.Time
}

/*
 *  Returns an instance of the Salesforce config - defaulting to the latest version of the API.
 *  Currently only supports UsernamePasswordAuthenticator.
 *  TODO: Create additional New functions to handle other Authenticaiton types.
 */
func New(instance, username, password, securityToken, consumerKey, consumerSecret string) Salesforce {
	auth := UsernamePasswordAuthenticator{
		Username:       username,
		Password:       password,
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		SecurityToken:  securityToken,
	}
	sf := Salesforce{
		InstanceUrl:    instance,
		HttpClient:     *http.DefaultClient,
		Authentication: auth,
		TokenValidFor:  2 * time.Hour, // Defaults to a 2 hour token expiry time
	}
	version, _ := sf.GetLatestVersion()
	sf.ApiVersionPath = version.Url
	return sf
}

func (sf *Salesforce) Authenticate() error {
	authResp, err := sf.Authentication.Authenticate(*sf)
	if err != nil {
		return err
	}
	sf.Token = authResp.AccessToken
	timestamp, err := strconv.ParseInt(authResp.IssuedAt, 10, 64)
	if err != nil {
		return err
	}
	seconds := timestamp / 1000
	t := time.Unix(seconds, 0)
	expiry := t.Add(sf.TokenValidFor)
	sf.TokenExpiryTime = expiry
	return nil
}
