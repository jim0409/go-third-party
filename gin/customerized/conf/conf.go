package conf

import "github.com/go-ini/ini"

var Conf *Config

type Config struct {
	BaseConf
	DbConf
	RedisConf
	LogConf
	WorkersPool
}

// BaseConf inlclude deatils server components
type BaseConf struct {
	HttpPort   string `ini:"HttpPort"` // http port
	Env        string `ini:"Env"`      // 運行環境
	ApiKey     string `ini:ApiKey`     // ApiKey
	SignKey    string `ini:SignKey`    // SingKey
	PubKeyPath string `ini:PubKeyPath` // 用於加密的Key的路徑
	PrvKeyPath string `ini:PrvKeyPath` // 用於解密的Key的路徑
	AesKey     string `ini:AesKey`     // (Default)用於Aes加密的字串
}

// DbConf is for mysql settings
type DbConf struct {
	DbName        string `ini:"DbName"`
	DbHost        string `ini:"DbHost"`
	DbPort        string `ini:"DbPort"`
	DbUser        string `ini:"DbUser"`
	DbPassword    string `ini:"DbPassword"`
	DbLogEnable   bool   `ini:"DbLogEnable"`
	DbMaxConnect  int    `ini:"DbMaxConnect"`
	DbIdleConnect int    `ini:"DbIdleConnect"`
}

// RedisConf pre-defined the redis conf include `addr/ auth/ idle/ active and timeout`
type RedisConf struct {
	RedisAddr        string `ini:"RedisAddr"`
	RedisAuth        string `ini:"RedisAuth"`
	RedisMaxIdle     int    `ini:"RedisMaxIdle"`
	RedisMaxActive   int    `ini:"RedisMaxActive"`
	RedisIdleTimeout int    `ini:"RedisIdleTimeout"`
}

// LogConf record log to specific folder
type LogConf struct {
	LogPath  string `ini:"LogPath"`
	LogLevel string `ini:"LogLevel"`
}

// WorkersPool declare workers pool
type WorkersPool struct {
	Tag        string `ini:"Tag"`
	NumWorkers int    `ini:"NumWorkers`
}

func InitConfig(confPath *string) (*Config, error) {
	Conf = new(Config)
	if err := ini.MapTo(Conf, *confPath); err != nil {
		return nil, err
	}
	return Conf, nil
}
