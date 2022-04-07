package config

type ServerConfig struct {
	Port        int         `mapstructure:"HOST_PORT"`
	MysqlConfig MysqlConfig `mapstructure:",squash"`
}

type MysqlConfig struct {
	Name     string `mapstructure:"MYSQL_DB_NAME"`
	Host     string `mapstructure:"MYSQL_HOST"`
	Port     int    `mapstructure:"MYSQL_PORT"`
	Username string `mapstructure:"MYSQL_USERNAME"`
	Password string `mapstructure:"MYSQL_PASSWORD"`
}
