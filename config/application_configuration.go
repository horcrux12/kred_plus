package config

type applicationConfiguration struct {
	Env string `json:"env" envconfig:"ENV"`
	App struct {
		Name    string `json:"name" envconfig:"APP_NAME"`
		Port    string `json:"port" envconfig:"APP_PORT"`
		Prefix  string `json:"prefix" envconfig:"APP_PREFIX"`
		Version string `json:"version" envconfig:"APP_VERSION"`
	} `json:"app"`
	Mysql struct {
		Host     string `json:"host" envconfig:"MYSQL_HOST"`
		Port     int    `json:"port" envconfig:"MYSQL_PORT"`
		Username string `json:"username" envconfig:"MYSQL_USERNAME"`
		Password string `json:"password" envconfig:"MYSQL_PASSWORD"`
		Database string `json:"database" envconfig:"MYSQL_DATABASE"`
	} `json:"mysql"`
	JWTSecretKey string `json:"jwt_secret_key" envconfig:"JWT_SECRET_KEY"`
}
