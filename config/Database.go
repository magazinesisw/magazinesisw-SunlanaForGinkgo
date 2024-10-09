package config

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"dataName"`
}

type DatabaseConfig struct {
	Database Database `yaml:"DataBase"`
}
