package goql

import (
	"sync"
)

type statement interface {
	exec(ctx *execCtx)
	explain(ctx *execCtx)
}

type execCtx struct {
	mu sync.RWMutex
}

type col struct {
	name       string
	constraint SQLConstraint
}

// createTableStmt represents a SQL Create Table statement.
type createTableStmt struct {
	exists    bool
	cols      []*col
	tableName string
}
