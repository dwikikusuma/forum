package configs

import (
	"github.com/spf13/viper"
	"log"
)

type option struct {
	configFolders []string
	configFile    string
	configType    string
}

type Option func(*option)

var config *Config

func Init(opts ...Option) error {
	opt := &option{
		configFolders: getDefaultFolders(),
		configFile:    getDefaultConfigFile(),
		configType:    getDefaultConfigType(),
	}

	for _, optFunc := range opts {
		optFunc(opt)
	}

	for _, folder := range opt.configFolders {
		viper.AddConfigPath(folder)
	}

	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to Read Config: %s", err.Error())
		return err
	}

	config = new(Config)
	return viper.Unmarshal(config)
}

func Get() *Config {
	if config == nil {
		return &Config{}
	}
	return config
}

func getDefaultFolders() []string {
	return []string{"./configs"}
}

func getDefaultConfigFile() string {
	return "config"
}

func getDefaultConfigType() string {
	return "yaml"
}

func WithConfigFolder(configFolders []string) Option {
	return func(opt *option) {
		opt.configFolders = configFolders
	}
}

func WithConfigFile(configFile string) Option {
	return func(opt *option) {
		opt.configFile = configFile
	}
}

func WithConfigType(configType string) Option {
	return func(opt *option) {
		opt.configType = configType
	}
}
