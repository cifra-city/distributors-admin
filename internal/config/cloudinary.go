package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

func InitCloudinaryClient(cfg Config) (*cloudinary.Cloudinary, error) {
	// URL connect format: cloudinary://<API_KEY>:<API_SECRET>@<CLOUD_NAME>
	cld, err := cloudinary.NewFromParams(cfg.Storage.CloudName, cfg.Storage.APIKey, cfg.Storage.APISecret)
	if err != nil {
		return nil, err
	}
	return cld, nil
}
