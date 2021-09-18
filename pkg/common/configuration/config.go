package configuration

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	API APIConfig

	//Storage
	Postgres PostgresConfig
}

type PostgresConfig struct {
	DatabaseName          string `envconfig:"DATABASE_NAME" default:"simple_bank_db"`
	User                  string `envconfig:"DATABASE_USER" default:"postgres"`
	Password              string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	Host                  string `envconfig:"DATABASE_HOST_DIRECT" default:"localhost"`
	Port                  string `envconfig:"DATABASE_PORT_DIRECT" default:"5432"`
	PoolMinSize           string `envconfig:"DATABASE_POOL_MIN_SIZE" default:"2"`
	PoolMaxSize           string `envconfig:"DATABASE_POOL_MAX_SIZE" default:"10"`
	PoolMaxConnLifetime   string `envconfig:"DATABASE_POOL_MAX_CONN_LIFETIME"`
	PoolMaxConnIdleTime   string `envconfig:"DATABASE_POOL_MAX_CONN_IDLE_TIME"`
	PoolHealthCheckPeriod string `envconfig:"DATABASE_POOL_HEALTHCHECK_PERIOD"`
	SSLMode               string `envconfig:"DATABASE_SSLMODE" default:"disable"`
	SSLRootCert           string `envconfig:"DATABASE_SSL_ROOTCERT"`
	SSLCert               string `envconfig:"DATABASE_SSL_CERT"`
	SSLKey                string `envconfig:"DATABASE_SSL_KEY"`
	Hostname              string `envconfig:"HOSTNAME" default:"localhost"`
}

func LoadConfig() (*Config, error) {
	var config Config
	noPrefix := ""

	err := envconfig.Process(noPrefix, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

type APIConfig struct {
	AppName string `envconfig:"APP_NAME" default:"simple-bank"`
	Port    string `envconfig:"API_PORT" default:"3000"`
}

func (c Config) DSN() string {
	connectString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		c.Postgres.User, c.Postgres.Password, c.Postgres.Host, c.Postgres.Port, c.Postgres.DatabaseName)

	if c.Postgres.SSLMode != "" {
		connectString = fmt.Sprintf("%s sslmode=%s",
			connectString, c.Postgres.SSLMode)
	}

	if c.Postgres.SSLRootCert != "" {
		connectString = fmt.Sprintf("%s sslrootcert=%s sslcert=%s sslkey=%s",
			connectString, c.Postgres.SSLRootCert, c.Postgres.SSLCert, c.Postgres.SSLKey)
	}

	if c.Postgres.Hostname != "" {
		connectString = fmt.Sprintf("%s application_name=%s", connectString, c.Postgres.Hostname)
	}

	// PoolMaxConnLifetime is the duration since creation after which a connection will be automatically closed.
	// default: 1h
	if c.Postgres.PoolMaxConnLifetime != "" {
		connectString = fmt.Sprintf("%s pool_max_conn_lifetime=%s", connectString, c.Postgres.PoolMaxConnLifetime)
	}

	// PoolMaxConnIdleTime is the duration after which an idle connection will be automatically closed by the health check.
	// default: 30m
	if c.Postgres.PoolMaxConnIdleTime != "" {
		connectString = fmt.Sprintf("%s pool_max_conn_idle_time=%s", connectString, c.Postgres.PoolMaxConnIdleTime)
	}

	// PoolHealthCheckPeriod is the duration between checks of the health of idle connections.
	// default: 1m
	if c.Postgres.PoolHealthCheckPeriod != "" {
		connectString = fmt.Sprintf("%s pool_healthcheck_period=%s", connectString, c.Postgres.PoolHealthCheckPeriod)
	}

	return connectString
}

func (c Config) URL() string {
	if c.Postgres.SSLMode == "" {
		c.Postgres.SSLMode = "disable"
	}

	connectString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Postgres.User, c.Postgres.Password, c.Postgres.Host, c.Postgres.Port, c.Postgres.DatabaseName, c.Postgres.SSLMode)

	if c.Postgres.SSLRootCert != "" {
		connectString = fmt.Sprintf("%s&sslrootcert=%s&sslcert=%s&sslkey=%s",
			connectString, c.Postgres.SSLRootCert, c.Postgres.SSLCert, c.Postgres.SSLKey)
	}

	return connectString
}
