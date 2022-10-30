// Copyright (c) 2022 Yandex LLC. All rights reserved.
// Author: Andrey Khaliullin <avhaliullin@yandex-team.ru>

package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"

	"go.uber.org/zap"
)

var globalLogger *zap.Logger

func initLogger() {
	config := zap.NewProductionConfig()
	config.DisableCaller = true
	level, err := zap.ParseAtomicLevel(c.LogLevel)
	if err != nil {
		level.SetLevel(zap.InfoLevel)
	}
	config.Level.SetLevel(level.Level())
	globalLogger, _ = config.Build()

	globalLogger = globalLogger.With(zap.String("instance_id", randomString(10)))
}

func withLogger(req *http.Request) *http.Request {
	requestID := req.URL.Query().Get("x-request-id")
	if requestID == "" {
		requestID = randomString(10)
	}
	return req.WithContext(ctxWithLog(req.Context(), globalLogger.With(zap.String("request-id", requestID))))
}

func randomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

type loggerKey struct{}

func ctxWithLog(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, log)
}

func logFromCtx(ctx context.Context) *zap.Logger {
	res := ctx.Value(loggerKey{})
	if res == nil {
		return globalLogger
	}
	return res.(*zap.Logger)
}
