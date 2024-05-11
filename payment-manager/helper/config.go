package helper

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName       string `mapstructure:"APPNAME"`
	Environment   string `mapstructure:"ENVIRONMENT"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	DBIdentity    string `mapstructure:"DB_IDENTITY"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpIn      string `mapstructure:"JWT_EXPIRED_IN"`
	TokenValidate string `mapstructure:"TOKEN_VALIDATE_URL"`
	BalanceSum    string `mapstructure:"BALANCE_SUM_URL"`
	BalanceReduc  string `mapstructure:"BALANCE_REDUC_URL"`
}

var MyConfig Config

func init() {
	Log.Println("Start Load Config")

	var err error
	if MyConfig, err = LoadConfig("."); err != nil {
		Log.Errorf("Load Config Failed : " + err.Error())
		panic(err)
	}
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	return
}
