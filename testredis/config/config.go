package config

import (
	"fmt"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name         string
	Info         Info
	QabircConfig QabircConfig
	ScaleMinute  bool
	AT           map[int]Atype
}
type Atype struct {
	ID   string
	Name string
}

type Info struct {
	Id int32
}

type QabircConfig struct {
	Redis Redis
}

type Redis struct {
	Address string
	Db      int
}

var logger *logging.Logger

var Globalconfig *Config

var defaultconfigpath = "./config/config.yaml"

//var defaultconfigpath = "./testredis/config/config.yaml"

func init() {
	logger = logging.MustGetLogger("logger")
}

func GetConfig() *Config {
	if Globalconfig != nil {
		return Globalconfig
	}
	Init("web", defaultconfigpath)
	return Globalconfig
}

func Init(envPrefix, path string) {
	Globalconfig = &Config{}
	if err := ParseConfig(envPrefix, path, Globalconfig); err != nil {
		fmt.Println(err)
		logger.Panic("Error parsing config file [%s] : %s", path, err)
	}
}

func ParseConfig(envPrefix, ymlFile string, object interface{}) error {
	v := newViper(envPrefix)
	v.SetConfigFile(ymlFile)
	err := v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read in config error: %s", err)
	}
	err = v.Unmarshal(object)
	if err != nil {
		return fmt.Errorf("unmarshal config to object error: %s", err)
	}
	return nil
}

func newViper(envPrefix string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	return v
}
