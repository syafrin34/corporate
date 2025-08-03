package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`

	JwtSecretKey string `json:"jwt_secret_key"`
	JwtIssuer    string `json:"jwt_issuer"`
}

type PsqlDB struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	DBName    string `json:"db_name"`
	DBMaxOpen int    `json:"db_max_open"`
	DBMaxIdle int    `json:"db_max_idle"`
}

type EmailConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Receiver string `json:"receiver"`
	IsTLS    bool   `json:"is_tls"`
}

type Config struct {
	App      App
	Psql     PsqlDB
	Supabase Supabase
	Email    EmailConfig
}
type Supabase struct {
	StorageUrl    string `json:"storage_url"`
	StorageKey    string `json:"storage_key"`
	StorageBucket string `json:"storage_bucket"`
}

func NewConfig() *Config {
	_ = godotenv.Load()
	viper.AutomaticEnv()
	fmt.Println("DB HOST:", viper.GetString("DATABASE_HOST"))
	return &Config{
		App: App{
			AppPort: viper.GetString("APP_PORT"),
			AppEnv:  viper.GetString("APP_PORT"),

			JwtSecretKey: viper.GetString("JWT_SECRET_KEY"),
			JwtIssuer:    viper.GetString("JWT_SECRET_KEY"),
		},
		Psql: PsqlDB{
			Host:      viper.GetString("DATABASE_HOST"),
			Port:      viper.GetString("DATABASE_PORT"),
			User:      viper.GetString("DATABASE_USER"),
			Password:  viper.GetString("DATABASE_PASSWORD"),
			DBName:    viper.GetString("DATABASE_NAME"),
			DBMaxOpen: viper.GetInt("DATABASE_MAX_OPEN_CONNECTION"),
			DBMaxIdle: viper.GetInt("DATABASE_MAX_IDLE_CONNECTION"),
		},
		Supabase: Supabase{
			StorageUrl:    viper.GetString("SUPABASE_STORAGE_URL"),
			StorageKey:    viper.GetString("SUPABASE_STORAGE_KEY"),
			StorageBucket: viper.GetString("SUPABASE_STORAGE_BUCKET"),
		},
		Email: EmailConfig{
			Host:     viper.GetString("EMAIL_HOST"),
			Port:     viper.GetInt("EMAIL_PORT"),
			Username: viper.GetString("EMAIL_USERNAME"),
			Password: viper.GetString("EMAIL_PASSWORD"),
			Receiver: viper.GetString("EMAIL_RECEIVER"),
			IsTLS:    viper.GetBool("EMAIL_TLS"),
		},
	}
}
