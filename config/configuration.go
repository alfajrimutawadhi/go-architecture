package config

type Config struct {
	DB          Database `json:"database"`
	Environment string   `json:"environment"`
}

type ShareConfig struct {
	DB          DBConfig
	Environment string
}

type Database struct {
	Development DBConfig `json:"development"`
	Production  DBConfig `json:"production"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	DBAesKey string `json:"db_aes_key"`
}
