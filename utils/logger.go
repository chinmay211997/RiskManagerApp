package utils

import (
	"log/slog"
	"os"

	"github.com/chinmay211997/RiskManagerApp/config"
)

var jsonHandler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
	Level: config.LogLevel,
})
var Logger = slog.New(jsonHandler)
