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
