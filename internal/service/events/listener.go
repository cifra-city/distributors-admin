package events

import (
	"context"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/distributors-admin/internal/config"
	"github.com/sirupsen/logrus"
)

func Listener(ctx context.Context) {
	_, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}
}
