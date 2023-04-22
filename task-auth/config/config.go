package config

import (
	"github.com/spf13/viper"
)

var (
	CookieName     string
	CookieMaxAge   int
	CookiePath     string
	CookieDomain   string
	CookieSecure   bool
	CookieHttpOnly bool

	Port     string
	DbName   string
	Host     string
	User     string
	Password string

	Domain  string
	AppPort string

	Origins string
)

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	//Cookie conf
	CookieName = viper.GetString("Cookie.name")
	CookieMaxAge = viper.GetInt("Cookie.maxAge")
	CookiePath = viper.GetString("Cookie.path")
	CookieDomain = viper.GetString("Cookie.domain")
	CookieSecure = viper.GetBool("Cookie.secure")
	CookieHttpOnly = viper.GetBool("Cookie.httpOnly")

	//Db conf
	Port = viper.GetString("DBserver.port")
	DbName = viper.GetString("DBserver.name")
	Host = viper.GetString("DBserver.host")
	User = viper.GetString("database.user")
	Password = viper.GetString("database.password")

	//App conf
	Domain = viper.GetString("AppServer.domain")
	AppPort = viper.GetString("AppServer.port")

	//CORS conf
	Origins = viper.GetString("CORS.origins")

	return nil
}
