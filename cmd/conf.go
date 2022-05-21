package main

type Config struct {
	Address  string
	Database struct {
		Driver string `yaml:"driver"`
		Dsn    string `yaml:"dsn"`
	}
	Redis struct {
		Hostname string `yaml:"hostname"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	}
	Gcp struct {
		ProjecId string `yaml:"project-id"`
	}
	BoxApi struct {
		Files  string `yaml:"files"`
		Upload string `yaml:"upload"`
	}
}

type BoxCredential struct {
	BoxAppSettings struct {
		ClientID     string `yaml:"clientID"`
		ClientSecret string `yaml:"clientSecret"`
	}
	AppAuth struct {
		PublicKeyID string `yaml:"publicKeyID"`
		Passphrase  string `yaml:"passphrase"`
	}
	EnterpriseID string `yaml:"enterpriseID"`
	UserId       string `yaml:"userId"`
}
