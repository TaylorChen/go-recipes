package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	ctx := context.WithValue(context.Background(), "reqID", "abc-123")
	logger.InfoContext(ctx, "starting", "module", "slog-demo", "version", "v1")
	logger.Error("something went wrong", "err", "boom", "code", 500)
}
