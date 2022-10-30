package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	initConfigFromEnv()
	initLogger()
	err := startServer()
	if err != nil {
		globalLogger.Error(fmt.Sprintf("server failed: %s", err), zap.Error(err))
	}
}
