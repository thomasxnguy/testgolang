package todofinder

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	Network   string
	ListenOn  string
	Root      string
	EnableTls bool
	KeyFile   string
	CertFile  string
}


//load server configuration file from path
//viper also have the ability to read from env variables
func LoadConfiguration(filepath *string) (*Configuration, error) {
	viper := viper.New()
	viper.SetConfigFile(*filepath)
	viper.SetConfigType("yaml")

	logrus.Printf("Loading configuration from %v", viper.ConfigFileUsed())
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	configuration := &Configuration{}
	viper.SetDefault("network", "tcp4")
	configuration.Network = viper.GetString("server.network")
	viper.SetDefault("listenOn", ":8080")
	configuration.ListenOn = viper.GetString("server.listenOn")
	viper.SetDefault("enableTls", "true")
	configuration.EnableTls = viper.GetBool("server.enableTls")
	configuration.EnableTls = viper.GetBool("server.keyFile")
	configuration.EnableTls = viper.GetBool("server.certFile")
	//default root directory
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	viper.SetDefault("root", dir)
	configuration.EnableTls = viper.GetBool("root")

	return configuration, nil
}
