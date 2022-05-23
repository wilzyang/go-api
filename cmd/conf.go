package main

type Config struct {
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`
	Database struct {
		Driver string `yaml:"driver"`
		Dsn    string `yaml:"dsn"`
	} `yaml:"database"`
	Redis struct {
		Hostname string `yaml:"hostname"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	} `yaml:"redis"`
	Gcp struct {
		CredPath string `yaml:"cred"`
		ProjecId string `yaml:"project-id"`
		Bucket   string `yaml:"bucket"`
	} `yaml:"gcp"`
	BoxApi struct {
		Files  string `yaml:"files"`
		Upload string `yaml:"upload"`
	} `yaml:"boxapi"`
}
