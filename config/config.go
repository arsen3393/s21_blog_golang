package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
	"time"
)

type Config struct {
	DbConfig     DataBaseConfig
	ServerConfig HttpServerConfig
	ClientConf   ClientConfig
}

type ClientConfig struct {
	Port            int    `env:"FRONT_PORT" envDefault:"8888"`
	StaticFilesPath string `env:"FRONT_STATIC_PATH" envDefault:"./web/"`
}

type HttpServerConfig struct {
	Port           string        `env:"SERVER_PORT" envDefault:"8080"`
	Timeout        time.Duration `env:"SERVER_TIMEOUT" envDefault:"5s"`
	IdleTimeout    time.Duration `env:"SERVER_IDLE_TIMEOUT" envDefault:"60s"`
	SwaggerEnabled bool          `env:"SERVER_SWAGGER_ENABLED" envDefault:"false"`
}

type DataBaseConfig struct {
	Host               string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port               string `env:"DB_PORT" envDefault:"5432"`
	Name               string `env:"DB_NAME" envDefault:"dev"`
	User               string `env:"DB_USER" envDefault:"root"`
	Password           string `env:"DB_PASSWORD" envDefault:"123456"`
	TimeZone           string `env:"DB_TIME_ZONE" env-default:"UTC" comment:"Часовой пояс базы данных"`
	MaxIdleConnections int    `env:"DB_MAX_IDLE_CONNECTIONS" envDefault:"40" comment:"Максимальное число простых соединений"`
	MaxOpenConnections int    `env:"DB_MAX_OPEN_CONNECTIONS" envDefault:"40" comment:"Максимальное число открытых соединений"`
}

func (d DataBaseConfig) ToString() string {
	return fmt.Sprintf("host = %s port = %s dbname = %s user = %s password = %s sslmode = disable", d.Host, d.Port, d.Name, d.User, d.Password)
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}
		if err := cleanenv.ReadConfig(".env", config); err != nil {
			log.Fatalf("Ошибка загрузки конфигурации: %v", err)
		}
	})
	return config
}
