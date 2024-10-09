package config

// Login 嵌套结构体
type Login struct {
	Phone    string `yaml:"Phone"`
	Password string `yaml:"Password"`
}

// LoginConfig 这里用yaml的顶级结构体
type LoginConfig struct {
	LoginInfo Login `yaml:"user"`
}
