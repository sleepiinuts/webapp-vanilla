package helpers

import (
	"log/slog"
	"os"
)

func Fatal(msg string, err error, logger *slog.Logger) {
	logger.Error(msg, "error", err)
	os.Exit(1)
}
