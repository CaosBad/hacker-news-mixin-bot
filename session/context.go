package session

import (
	"context"
	"database/sql"

	"github.com/crossle/hacker-news-mixin-bot/durable"
)

type contextValueKey int // custom type of int 

const (
	keyRequest  contextValueKey = 0
	keyDatabase contextValueKey = 1
	keyLogger   contextValueKey = 2
)
// get logger from context
func Logger(ctx context.Context) *durable.Logger {
	v, _ := ctx.Value(keyLogger).(*durable.Logger)
	return v
}
// get db from context
func Database(ctx context.Context) *sql.DB {
	v, _ := ctx.Value(keyDatabase).(*sql.DB)
	return v
}
// inject logger into context
func WithLogger(ctx context.Context, logger *durable.Logger) context.Context {
	return context.WithValue(ctx, keyLogger, logger)
}
// inject db into context
func WithDatabase(ctx context.Context, database *sql.DB) context.Context {
	return context.WithValue(ctx, keyDatabase, database)
}
