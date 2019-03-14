package logger

import (
	"cloud.google.com/go/logging"
	"context"
)

func NewLogger() *logging.Logger {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, "projects/turnkey-env-234412")
	if err != nil {
		panic(err)
	}
	return client.Logger("syslog")
}
