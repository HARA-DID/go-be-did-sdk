package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/meQlause/go-be-did/internal/infrastructure/sdk"
)

var (
	config *Config
)

type Config struct {
	App AppConfig
	HNS HNSConfig
}

type AppConfig struct {
	Port       string
	Version    string
	LogLevel   string
	RPCTimeout int
}

type HNSConfig struct {
	AccountAbstraction sdk.AccountAbstractionHNS
	// DIDRoot            sdk.DIDRootHNS
}

func InitConfig() {
	once.Do(func() {
		cfg, err := Load()
		if err != nil {
			panic(err)
		}
		config = cfg
	})
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		App: AppConfig{
			Port:       getEnv("PORT", "8080"),
			LogLevel:   getEnv("LOG_LEVEL", "info"),
			Version:    getEnv("VERSION", "1.0"),
			RPCTimeout: getEnvAsInt("RPC_TIMEOUT", 30),
		},
		HNS: HNSConfig{
			AccountAbstraction: sdk.AccountAbstractionHNS{
				EntryPoint:    getEnv("AA_ENTRYPOINT_HNS", ""),
				GasManager:    getEnv("AA_GAS_MANAGER_HNS", ""),
				Wallet:        getEnv("AA_WALLET_HNS", ""),
				WalletFactory: getEnv("AA_WALLET_FACTORY_HNS", ""),
			},
			// DIDRoot: sdk.DIDRootHNS{
			// 	RootFactory: getEnv("DID_ROOT_FACTORY_HNS", ""),
			// 	RootStorage: getEnv("DID_ROOT_STORAGE_HNS", ""),
			// },
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	if c.App.Port == "" {
		return errors.New("PORT is required")
	}

	if err := c.HNS.AccountAbstraction.Validate(); err != nil {
		return err
	}

	// if err := c.HNS.DIDRoot.Validate(); err != nil {
	// 	return err
	// }

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}

func GetApp() AppConfig {
	return config.App
}

func GetConfig() *Config {
	return config
}
