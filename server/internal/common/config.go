package common

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	DatabaseType           string
	PostgresConfig         PostgresDatabaseConfig
	BucketConnectionConfig BucketConfig
}

func LoadConfig() Config {
	dbType := GetEnv("DB_TYPE", "postgres")

	config := &Config{
		DatabaseType: GetEnv("DB_TYPE", "postgres"),
	}

	logrus.WithField("type", dbType).Info("loading database config")
	return *config
}

type PostgresDatabaseConfig struct {
	ConnectionString string
}

func (c *Config) LoadPostgresDatabaseConfig() PostgresDatabaseConfig {

	return PostgresDatabaseConfig{
		ConnectionString: GetEnv("DATABASE_URL", "postgresql://postgres:postgres@localhost"),
	}
}

type Neo4jConfig struct {
	URI      string
	Username string
	Password string
}

func (c *Config) LoadNeo4jConfig() Neo4jConfig {

	return Neo4jConfig{
		URI:      GetEnv("GRAPH_URL", "bolt://graph-db:7687"),
		Username: GetEnv("GRAPH_USERNAME", "neo4j"),
		Password: GetEnv("GRAPH_PASSWORD", "test"),
	}
}

type BucketConfig struct {
	ConnectionString string
	AccessID         string
	AccessKey        string
}

func (c *Config) LoadBucketConfig() BucketConfig {
	host := GetEnv("BUCKET_HOST", "s3://my-books")
	accessID := os.Getenv("ACCESS_ID")
	key := os.Getenv("ACCESS_KEY")
	return BucketConfig{
		ConnectionString: host,
		AccessID:         accessID,
		AccessKey:        key,
	}
}

func GetEnv(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
