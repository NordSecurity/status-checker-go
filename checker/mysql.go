package checker

import (
	"context"
	"database/sql"
	"time"

	"github.com/NordSec/status-checker-go"
)

const timeout = time.Second

func NewMysqlChecker(db *sql.DB) status.Checker {
	return &mysqlChecker{db}
}

type mysqlChecker struct {
	db *sql.DB
}

func (mc mysqlChecker) Name() string {
	return "mysql"
}

func (mc mysqlChecker) Status() status.Status {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()

	if err := mc.db.PingContext(ctx); err != nil {
		return status.DOWN
	}

	return status.OK
}
