### GoSF - Work In Progress

A go library for interacting with the Salesforce APIs and a command line tool to interact with Salesforce.


#### Usage

`gosf query "SELECT Id FROM Account"`

#### Configuration

Viper is used to load in configuration values from either Environment Variables or a .env file.

```
SF_INSTANCE=https://login.salesforce.com
SF_USERNAME=user@user.com
SF_PASSWORD=PASSWORD
SF_SECURITY_TOKEN=ABCXYZ123
SF_CONSUMER_KEY=ZYXCBA321
SF_CONSUMER_SECRET=987FGHLMN
```

#### Authentication
 - [OAuth 2.0 Username-Password Flow](https://help.salesforce.com/s/articleView?id=sf.remoteaccess_oauth_username_password_flow.htm&type=5)
    - Consumer Key
    - Consumer Secret
    - Username
    - Password
    - Security Token
    - Under Setup -> Identity -> OAuth and OpenID Connect Settings - Enable "Allow OAuth Username-Password Flows"
 - [OAuth 2.0 Client Credentials](https://help.salesforce.com/s/articleView?id=sf.remoteaccess_oauth_client_credentials_flow.htm&type=5)
    - Consumer Key
    - Consumer Secret
    - Setup Steps
      - Ensure Client Credentials flow is enabled when setting up the Connected App.
      - Add the `api` scope to the Connected App. [Link](https://trailhead.salesforce.com/trailblazer-community/feed/0D54V00007T4FmsSAF).
      - Once created go to Manage Connected Apps & Edit Policies of the App then set the [Client Credentials Flow -> Run As User](https://trailhead.salesforce.com/trailblazer-community/feed/0D54V00007T4L8NSAV)
      - The user should have the API Only permission.
  
#### To Do
 - Additional Authentication Methods
  - [OAuth 2.0 JWT Bearer](https://help.salesforce.com/s/articleView?id=sf.remoteaccess_oauth_jwt_flow.htm&type=5)
  - [OAuth 2.0 Device Flow](https://help.salesforce.com/s/articleView?id=sf.remoteaccess_oauth_device_flow.htm&type=5)
 - Handle commands e.g. query, subscribe, publish, insert, update, upsert (?)
 - Pub/Sub Events
   - [Pub/Sub API as a gRPC API](https://developer.salesforce.com/docs/platform/pub-sub-api/guide/grpc-api.html)
   - [Go Example](https://github.com/forcedotcom/pub-sub-api/tree/main/go)

