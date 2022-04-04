package config

import (
	"ecos/client/credentials"
	"ecos/utils/config"
	"ecos/utils/logger"
	"os"
	"time"
)

const chunkSize = 1 << 20
const uploadTimeout = time.Second * 10
const uploadBuffer = 1 << 25

type ObjectConfig struct {
	ChunkSize uint64
}

type ClientConfig struct {
	config.Config
	Credential    credentials.Credential
	Object        ObjectConfig
	UploadTimeout time.Duration
	UploadBuffer  uint64
	NodeAddr      string
	NodePort      uint64
}

var DefaultConfig *ClientConfig
var Config *ClientConfig

func init() {
	if DefaultConfig == nil {
		DefaultConfig = &ClientConfig{
			Config:     config.Config{},
			Credential: credentials.New("root", "root"),
			Object: ObjectConfig{
				ChunkSize: chunkSize,
			},
			UploadTimeout: uploadTimeout,
			UploadBuffer:  uploadBuffer,
			NodeAddr:      "ecos-edge-dev",
			NodePort:      3267,
		}
	}
	if Config == nil {
		err := InitConfig(Config)
		if err != nil {
			logger.Warningf("parse client config failed %v", err)
			logger.Warningf("using default client config")
			Config = DefaultConfig
		}
	}
}

// InitConfig check config and init data dir and set some empty config value
func InitConfig(conf *ClientConfig) error {
	// read persist config file in storage path
	// TODO: read confPath from cmd args
	confPath := "./config/client.json"
	s, err := os.Stat(confPath)
	if err == nil && !s.IsDir() && s.Size() > 0 {
		config.Register(DefaultConfig, confPath)
		config.ReadAll()
	}
	err = config.GetConf(conf)
	return err
}
