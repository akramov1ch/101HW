package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    ServerAddr  string
    RedisAddr   string
    WorkerCount int
    QueueSize   int
}

func LoadConfig() (*Config, error) {
    viper.SetConfigFile(".env")
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    config := &Config{
        ServerAddr:  viper.GetString("SERVER_ADDR"),
        RedisAddr:   viper.GetString("REDIS_ADDR"),
        WorkerCount: viper.GetInt("WORKER_COUNT"),
        QueueSize:   viper.GetInt("QUEUE_SIZE"),
    }

    return config, nil
}
