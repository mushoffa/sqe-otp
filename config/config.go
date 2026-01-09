package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	BASE_API = "/api/v1"
)

type Config struct {
	Database   Database   `mapstructure:"db"`
	HttpServer HttpServer `mapstructure:"server"`
	Redis      Redis      `mapstructure:"redis"`
}

type Database struct {
	Driver      string   `mapstructure:"driver"`
	Host        string   `mapstructure:"host"`
	Replica     []string `mapstructure:"replica"`
	Port        int      `mapstructure:"port"`
	Name        string   `mapstructure:"name"`
	Schema      string   `mapstructure:"schema"`
	SSL         string   `mapstructure:"ssl"`
	Username    string   `mapstructure:"username"`
	Password    string   `mapstructure:"password"`
	MaxOpenConn int      `mapstructure:"max_open_conn"`
	MaxIdleConn int      `mapstructure:"max_idle_conn"`
	MaxLifeTime int      `mapstructure:"max_life_time"`
	MaxIdleTime int      `mapstructure:"max_idle_time"`
}

func (db Database) Dsn(host string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d search_path=%s sslmode=disable", host, db.Username, db.Password, db.Name, db.Port, db.Schema)
}

func (db Database) PrimaryDsn() string {
	return db.Dsn(db.Host)
}

type HttpServer struct {
	Port      int `mapstructure:"port"`
	JwtSecret []byte
}

type Redis struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (c Redis) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func Get() *Config {
	c := new(Config)
	viper.AddConfigPath("./config")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(c); err != nil {
		panic(err)
	}

	return c
}
