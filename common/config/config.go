package config

type Config struct {
	Mysql
	Redis
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
