package config

import (
	"github.com/cifra-city/cifra-rabbit"
	"github.com/cifra-city/distributors-admin/internal/data/sql"
	"github.com/cifra-city/tokens"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/sirupsen/logrus"
)

const (
	SERVER = "server"
)

type Service struct {
	Config       *Config
	SqlDB        *sql.Repo
	Logger       *logrus.Logger
	Cloud        *cloudinary.Cloudinary
	TokenManager *tokens.TokenManager
	Broker       *cifra_rabbit.Broker
}

func NewServer(cfg *Config) (*Service, error) {
	logger := SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	queries, err := sql.NewRepoSQL(cfg.Database.URL)
	if err != nil {
		return nil, err
	}
	TokenManager := tokens.NewTokenManager(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB, logger, cfg.JWT.AccessToken.TokenLifetime)
	broker, err := cifra_rabbit.NewBroker(cfg.Rabbit.URL, cfg.Rabbit.Exchange)
	if err != nil {
		return nil, err
	}
	Cloud, err := InitCloudinaryClient(*cfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		Config:       cfg,
		SqlDB:        queries,
		Logger:       logger,
		TokenManager: TokenManager,
		Cloud:        Cloud,
		Broker:       broker,
	}, nil
}
