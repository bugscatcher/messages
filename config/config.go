package config

import (
	"github.com/bugscatcher/messages/grpcHelper"
	"github.com/bugscatcher/messages/kafka"
	"github.com/bugscatcher/messages/postgresql"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	PostgreSQL       postgresql.Config     `mapstructure:"postgresql"`
	PublicGRPCServer grpcHelper.ServerConf `mapstructure:"public_grpc_server"`
	KafkaConf        kafka.Config          `mapstructure:"kafka"`
}

func New() (Config, error) {
	var conf Config
	err := viper.Unmarshal(&conf)
	return conf, err
}

func init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	_ = err

	// PostgreSQL
	viper.SetDefault("postgresql.port", 5432)
	viper.SetDefault("postgresql.host", "localhost")
	viper.SetDefault("postgresql.database", "messages")
	viper.SetDefault("postgresql.user", "postgres")
	viper.SetDefault("postgresql.password", "")
	viper.SetDefault("postgresql.secured", false)
	viper.SetDefault("postgresql.max_connections", 10)
	viper.SetDefault("postgresql.max_connection_age", time.Minute)

	// GRPC server
	viper.SetDefault("public_grpc_server.port", 10001)
	viper.SetDefault("public_grpc_server.address", "0.0.0.0")

	// Kafka
	viper.SetDefault("kafka.brokers", []string{"kafka:9092"})
}
