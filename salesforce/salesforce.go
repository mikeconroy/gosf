package salesforce

type Salesforce struct {
	InstanceUrl    string
	Username       string
	Password       string
	ConsumerKey    string
	ConsumerSecret string
	SecurityToken  string
	ApiVersionPath string
}

/*
 *  Returns an instance of the Salesforce config - defaulting to the latest version of the API.
 */
func New(instance, username, password, securityToken, consumerKey, consumerSecret string) Salesforce {
	sf := Salesforce{
		InstanceUrl:    instance,
		Username:       username,
		Password:       password,
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		SecurityToken:  securityToken,
	}
	version, _ := sf.GetLatestVersion()
	sf.ApiVersionPath = version.Url
	return sf
}

func (sf Salesforce) Authenticate() (string, error) {

	return "", nil
}
