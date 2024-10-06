package config

import (
	"log/slog"
	"time"
)

const AppName = "RiskHealthApp"

const HOST = "127.0.0.1"
const PORT = "8080"

const ReadTimeout = 1 * time.Second
const WriteTimeout = 1 * time.Second

const LogLevel = slog.LevelDebug
