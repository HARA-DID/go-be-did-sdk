package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	aasdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
	dasdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didalias"
	drsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didroot"
	// didvcsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didvc"
)

var (
	config *Config
)

type Config struct {
	App AppConfig
	HNS HNSConfig
	RPC RPCConfig
}

type RPCConfig struct {
	Endpoints []string
}

type AppConfig struct {
	Port       string
	Version    string
	LogLevel   string
	RPCTimeout int
}

type HNSConfig struct {
	AccountAbstraction aasdk.AccountAbstractionHNS
	DIDRoot            drsdk.DIDRootHNS
	// DIDVC              didvcsdk.DIDVCHNS
	DIDAlias dasdk.DIDAliasHNS
}

func InitConfig() {
	onceConfig.Do(func() {
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
			AccountAbstraction: aasdk.AccountAbstractionHNS{
				EntryPoint:    getEnv("AA_ENTRYPOINT_HNS", ""),
				GasManager:    getEnv("AA_GAS_MANAGER_HNS", ""),
				Wallet:        getEnv("AA_WALLET_HNS", ""),
				WalletFactory: getEnv("AA_WALLET_FACTORY_HNS", ""),
			},
			DIDRoot: drsdk.DIDRootHNS{
				RootFactory: getEnv("DID_ROOT_FACTORY_HNS", ""),
				RootStorage: getEnv("DID_ROOT_STORAGE_HNS", ""),
			},
			// DIDVC: didvcsdk.DIDVCHNS{
			// 	VCFactory:      getEnv("DID_VC_FACTORY_HNS", ""),
			// 	VCStorage:      getEnv("DID_VC_STORAGE_HNS", ""),
			// 	CertificateNFT: getEnv("DID_VC_CERTIFICATE_NFT_HNS", ""),
			// 	IdentityNFT:    getEnv("DID_VC_IDENTITY_NFT_HNS", ""),
			// },
			DIDAlias: dasdk.DIDAliasHNS{
				AliasFactory: getEnv("DID_ALIAS_FACTORY_HNS", ""),
				AliasStorage: getEnv("DID_ALIAS_STORAGE_HNS", ""),
			},
		},
		RPC: RPCConfig{
			Endpoints: getRPCs(),
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func getRPCs() []string {
	var rpcs []string
	for i := 1; i <= 4; i++ {
		key := "RPC" + strconv.Itoa(i)
		if v := os.Getenv(key); v != "" {
			rpcs = append(rpcs, v)
		}
	}
	return rpcs
}

func (c *Config) Validate() error {
	if c.App.Port == "" {
		return errors.New("PORT is required")
	}

	if len(c.RPC.Endpoints) == 0 {
		return errors.New("at least one RPC endpoint is required")
	}

	if err := c.HNS.AccountAbstraction.Validate(); err != nil {
		return err
	}

	if err := c.HNS.DIDRoot.Validate(); err != nil {
		return err
	}

	// if err := c.HNS.DIDVC.Validate(); err != nil {
	// 	return err
	// }

	if err := c.HNS.DIDAlias.Validate(); err != nil {
		return err
	}

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
