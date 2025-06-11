package config

type Config struct {
	Mysql
	Redis
	Server
	Elastic
	AliyunGpt
	Mongo
}

type Server struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type Elastic struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
	User string `yaml:"User"`
	Pass string `yaml:"Pass"`
}

type Mysql struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	UserName string `yaml:"UserName"`
	Password string `yaml:"PassWord"`
	Database string `yaml:"DataBase"`
}

type Redis struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Prot"`
	User     string `yaml:"User"`
	Password string `yaml:"PassWord"`
	DB       int    `yaml:"DB"`
}

type AliyunGpt struct {
	ApiKey string `yaml:"ApiKey"`
}

type Mongo struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
	User string `yaml:"User"`
	Pass string `yaml:"Pass"`
}
