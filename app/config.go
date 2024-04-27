package app

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Global  Global  `toml:"global"`
	Bitcoin Bitcoin `toml:"bitcoin"`
	Side    Side    `toml:"side"`
}

type Global struct {
	LogLevel string `toml:"log-level"                   comment:"log level of the daemon"`
}

type Bitcoin struct {
	// Bitcoin specific configuration
	RPC         string `toml:"rpc"                           comment:"Bitcoin RPC endpoint"`
	RPCUser     string `toml:"rpcuser"                   comment:"Bitcoin RPC user"`
	RPCPassword string `toml:"rpcpassword"           comment:"Bitcoin RPC password"`

	Frequency int    `toml:"frequency"                  comment:"frequency of Bitcoin block polling in seconds"`
	Sender    string `toml:"sender"                     comment:"Bitcoin sender address"`
}

type Side struct {
	// Side specific configuration
	RPC  string `toml:"rpc"                           comment:"Side RPC endpoint"`
	REST string `toml:"rest"                          comment:"Side REST endpoint"`

	Frequency int    `toml:"frequency"                  comment:"frequency of Side block polling in	seconds"`
	Sender    string `toml:"sender"                     comment:"Side sender address"`
}

func defaultConfig() *Config {
	return &Config{
		Global: Global{
			LogLevel: "info",
		},
		Bitcoin: Bitcoin{
			RPC:         "localhost:8332",
			RPCUser:     "side",
			RPCPassword: "12345678",
			Frequency:   10 * 60 * 60,
			Sender:      "",
		},
		Side: Side{
			RPC:       "http://localhost:26657",
			REST:      "http://localhost:1317",
			Frequency: 6,
			Sender:    "",
		},
	}
}

const AppName = "shuttler"

var DefaultHome = filepath.Join(os.Getenv("HOME"), ".shuttler")

var DefaultConfigFilePath = DefaultHome + "/config/config.toml"

type ConfigBuilder struct {
	homePath string
}

func NewConfigBuilder(homePath string) *ConfigBuilder {
	if homePath == "" {
		homePath = DefaultHome
	}
	return &ConfigBuilder{
		homePath: homePath,
	}
}

func (c *ConfigBuilder) ConfigFilePath() string {
	return c.homePath + "/config/config.toml"
}

func (c *ConfigBuilder) InitConfig() *Config {
	cfg := defaultConfig()
	out, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	os.MkdirAll(c.homePath+"/config", 0755)

	err = os.WriteFile(c.ConfigFilePath(), out, 0644)
	if err != nil {
		panic(err)
	}
	return cfg
}

func (c *ConfigBuilder) LoadConfigFile() *Config {
	in, err := os.ReadFile(c.ConfigFilePath())
	if err != nil {
		panic(err)
	}
	cfg := &Config{}
	err = toml.Unmarshal(in, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

// func (c *ConfigBuilder) RuntimeConfig(ctx context.Context, a *appState) (*Config, error) {
// 	return c, nil
// }

func (c *ConfigBuilder) validateConfig() error {
	// validate config
	return nil
}
