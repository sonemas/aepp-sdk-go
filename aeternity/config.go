package aeternity

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	utils "github.com/aeternity/aepp-sdk-go/utils"

	yaml "gopkg.in/yaml.v2"
)

const (
	ConfigFilename = "aecli.config"
)

// EpochConfig configuration for the epoch node
type EpochConfig struct {
	URL          string `yaml:"url" mapstructure:"url"`
	InternalURL  string `yaml:"internal_url" mapstructure:"internal_url"`
	WebsocketURL string `yaml:"websocket_url" mapstructure:"websocket_url"`
}

type AetNameConfig struct {
	TTL       int64 `yaml:"ttl" mapstructure:"ttl"`
	ClientTTL int64 `yaml:"client_ttl" mapstructure:"client_ttl"`
}

type AetContractConfig struct {
	Gas       int64 `yaml:"gas" mapstructure:"gas"`
	GasPrice  int64 `yaml:"gas_price" mapstructure:"gas_price"`
	Deposit   int64 `yaml:"deposit" mapstructure:"deposit"`
	VMVersion int64 `yaml:"vm_version" mapstructure:"vm_version"`
}

// ClientConfig client paramters configuration
type ClientConfig struct {
	TxTTL      int64             `yaml:"tx_ttl" mapstructure:"tx_ttl"`
	Fee        int64             `yaml:"fee" mapstructure:"fee"`
	DefaultKey string            `yaml:"default_key_name" mapstructure:"default_key_name"`
	Names      AetNameConfig     `yaml:"names" mapstructure:"names"`
	Contracts  AetContractConfig `yaml:"contracts" mapstructure:"contracts"`
}

// TuningConfig fine tuning of parameters of the client
type TuningConfig struct {
	ChainPollInteval float32 `yaml:"chain_poll_interval" mapstructure:"chain_poll_interval"`
	TxPayload        string  `yaml:"tx_payload" mapstructure:"tx_payload"`
	ResponseEncoding string  `yaml:"msg_encoding" mapstructure:"msg_encoding"`
}

// ProfileConfig a configuration profile
type ProfileConfig struct {
	Name   string       `yaml:"name" mapstructure:"name"`
	Epoch  EpochConfig  `yaml:"epoch" mapstructure:"epoch"`
	Client ClientConfig `yaml:"client" mapstructure:"client"`
	Tuning TuningConfig `yaml:"tuning" mapstructure:"tuning"`
}

// ConfigSchema define the configuration object
type ConfigSchema struct {
	ActiveProfile   string           `yaml:"active_profile" mapstructure:"active_profile"`
	DefaultKeysPath string           `yaml:"keys_path" mapstructure:"keys_path"`
	Profiles        []*ProfileConfig `yaml:"profiles" mapstructure:"profiles"`
	P               *ProfileConfig   `yaml:"-" mapstructure:"-"` // holds the active profile
	ConfigPath      string           `yaml:"-" mapstructure:"-"` // the path of the configuration file
}

//Defaults generate configuration defaults
func (c *ProfileConfig) Defaults() *ProfileConfig {
	c.Client = ClientConfig{
		Contracts: AetContractConfig{},
		Names:     AetNameConfig{},
	}
	// for server
	utils.DefaultIfEmptyStr(&c.Epoch.URL, "https://sdk-testnet.aepps.com")
	utils.DefaultIfEmptyStr(&c.Epoch.InternalURL, "http://127.0.0.1:3113")
	utils.DefaultIfEmptyStr(&c.Epoch.WebsocketURL, "ws://127.0.0.1:3013")
	// for client
	utils.DefaultIfEmptyStr(&c.Client.DefaultKey, "wallet.key")
	utils.DefaultIfEmptyInt64(&c.Client.TxTTL, 500)
	utils.DefaultIfEmptyInt64(&c.Client.Fee, 1)
	utils.DefaultIfEmptyInt64(&c.Client.Names.TTL, 500)
	utils.DefaultIfEmptyInt64(&c.Client.Contracts.Gas, 40000000)
	utils.DefaultIfEmptyInt64(&c.Client.Contracts.GasPrice, 1)
	// for tuning
	utils.DefaultIfEmptyStr(&c.Tuning.TxPayload, "payload")
	utils.DefaultIfEmptyStr(&c.Tuning.ResponseEncoding, "json")
	return c
}

//Validate configuration
func (c *ProfileConfig) Validate() {

}

//Validate configuration
func (c *ConfigSchema) Validate() {
	valid := false
	for _, p := range c.Profiles {
		if p.Name == c.ActiveProfile {
			valid = true
		}
	}

	if !valid {
		fmt.Println("Invalid configuration")
		os.Exit(1)
	}
}

//Defaults configuration
func (c *ConfigSchema) Defaults() *ConfigSchema {
	for _, p := range c.Profiles {
		p.Defaults()
	}
	c.ActivateProfile(c.ActiveProfile)
	return c
}

// Config sytem configuration
var Config ConfigSchema

// GenerateDefaultConfig generate a default configuration
func GenerateDefaultConfig(outFile, version string) {
	Config := ConfigSchema{
		DefaultKeysPath: "keys",
		ConfigPath:      outFile,
	}
	Config.NewProfile("default")
	Config.ActivateProfile("default")
}

// Save save the configuration to disk
func (c *ConfigSchema) Save() {
	b, _ := yaml.Marshal(c)
	data := strings.Join([]string{
		"#",
		"#\n# Configuration for aepp-sdk-go \n#\n",
		fmt.Sprintf("#\n# update on %s \n#\n", time.Now().Format(time.RFC3339)),
		fmt.Sprintf("%s", b),
		"#\n# Config end\n#",
	}, "\n")
	err := ioutil.WriteFile(c.ConfigPath, []byte(data), 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("config file written to", c.ConfigPath)
}

// NewProfile create a new configuration profile
func (c *ConfigSchema) NewProfile(name string) {
	p := &ProfileConfig{
		Name: name,
	}
	p.Defaults()
	c.Profiles = append(c.Profiles, p)
}

// ActivateProfile a profile, err if the profile doesnt exists
func (c *ConfigSchema) ActivateProfile(name string) (err error) {
	p, err := c.GetProfile(name)
	if err != nil {
		return
	}
	c.ActiveProfile = name
	c.P = p
	return
}

// GetProfile get a profile by name, err when doenst exists
func (c *ConfigSchema) GetProfile(name string) (p *ProfileConfig, err error) {
	for _, p = range c.Profiles {
		if p.Name == name {
			return
		}
	}
	err = fmt.Errorf("Profile %s not found", name)
	return
}
