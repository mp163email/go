package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"sync"
	"time"
)

/*
配置类
*/
type Config struct {
	//服务器相关配置
	Server struct {
		Host         string        `yaml:"host"`
		Port         int           `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
		IdeleTimeout time.Duration `yaml:"idle_timeout"`
	} `yaml:"server"`

	//游戏相关配置
	Game struct {
		HeartbeanInterval time.Duration `yaml:"heartbeat_interval"` //最大多少秒没有收到心跳包就断开连接
		MaxPlayers        int           `yaml:"max_players"`
	} `yaml:"lit-game"`
}

// 定义一个包级别的全局变量，用于存储配置信息
// 而且在整个应用程序的生命周期里只会存在一个实例,相当于单例
var GlobalConfig Config
var loadOnce sync.Once

func Load(path string) (*Config, error) {
	var err error
	loadOnce.Do(func() { //使用sync.Once保证只加载一次（配置相关的就应该执行一次）
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			err = readErr
			return
		}
		var cfg Config
		unmarshalErr := yaml.Unmarshal(data, &cfg)
		if unmarshalErr != nil {
			err = unmarshalErr
			return
		}
		GlobalConfig = cfg
	})
	return &GlobalConfig, err
}
