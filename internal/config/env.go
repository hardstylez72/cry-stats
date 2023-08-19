package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Database struct {
		Conn string
	}

	ENV string

	Config struct {
		Env                ENV
		GRPCAddr           string
		DBConn             string
		MigrationsDir      string
		PrometheusPort     string
		JaegerUrl          string
		JaegerServiceName  string
		ArbiscanToken      string
		ARBIUSDTWalletAddr string
		Standalone         bool
	}
)

const (
	Dev   ENV = "dev"
	Local ENV = "local"
	Prod  ENV = "prod"
)

func envFromString(s string) ENV {
	switch s {
	case "dev":
		return Dev
	case "local":
		return Local
	case "prod":
		return Prod
	}
	panic("invalid env value: " + s)
}

var CFG *Config

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("can not load .env file, using os env")
	}
	c := &Config{
		Env:           envFromString(mustenv("ENV")),
		DBConn:        mustenv("POSTGRES_CONN"),
		GRPCAddr:      mustenv("GRPC_ADDR"),
		MigrationsDir: mayenv("MIGRATIONS_DIR", "/migrations"),

		PrometheusPort:    mayenv("PROMETHEUS_PORT", ""),
		JaegerServiceName: mayenv("JAEGER_SERVICE_NAME", ""),
		JaegerUrl:         mayenv("JAEGER_URL", ""),
		Standalone:        mayenv("STANDALONE", "false") == "true",
	}

	if !c.Standalone {
		c.ARBIUSDTWalletAddr = mustenv("ARBI_USDT_WALLET_ADDR")
		c.ArbiscanToken = mustenv("APIKEY_ARBISCAN")
	}

	CFG = c

	return c, nil
}
func mayenv(key string, dafaultVallue string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return dafaultVallue
}
func mustenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(key + " is not set")
	}
	return v
}
