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
		Bucket   string `yaml:"bucket"`
	}
	BoxApi struct {
		Files  string `yaml:"files"`
		Upload string `yaml:"upload"`
	}
}
