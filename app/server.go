package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func startServer() error {
	port := c.Port
	if port <= 0 {
		port = 80
	}
	globalLogger.Debug("starting server")
	return http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(handle))
}

func handle(rw http.ResponseWriter, req *http.Request) {
	defer globalLogger.Sync()

	req = withLogger(req)

	start := time.Now()
	res, err := doHandle(req)
	finish := time.Now()
	if err != nil {
		handleError(req.Context(), rw, err)
	} else {
		logFromCtx(req.Context()).Info(fmt.Sprintf("request processed in %v", finish.Sub(start)))
		rw.WriteHeader(200)
		writeJSON(req.Context(), rw, res)
	}
}

func handleError(ctx context.Context, rw http.ResponseWriter, err error) {
	log := logFromCtx(ctx)
	var uErr *UserError
	if errors.As(err, &uErr) {
		log.Warn(fmt.Sprintf("user error: %s", err))
		rw.WriteHeader(uErr.GetHTTPCode())
		writeJSON(ctx, rw, map[string]string{"error": err.Error()})
	} else {
		log.Error(fmt.Sprintf("internal error: %s", err), zap.Error(err))
		rw.WriteHeader(http.StatusInternalServerError)
		writeJSON(ctx, rw, map[string]string{"error": "internal error"})
	}
}

func writeJSON(ctx context.Context, rw http.ResponseWriter, response interface{}) {
	log := logFromCtx(ctx)
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Error(fmt.Sprintf("failed to marshal response: %s", err), zap.Error(err))
		return
	}
	_, err = rw.Write(bytes)
	if err != nil {
		log.Warn(fmt.Sprintf("failed to write response: %s", err), zap.Error(err))
	}
}
