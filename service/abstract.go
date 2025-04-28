package service

import (
	"context"
	"kredi-plus.com/be/app"
	"kredi-plus.com/be/lib/constanta"
	"kredi-plus.com/be/lib/exception"
	"strings"
)

func errorHandlerForConstraint(err error) (res error) {
	switch errStr := err.Error(); {
	case strings.Contains(errStr, "uq_user_username"):
		res = exception.AlreadyExist("username")
	case strings.Contains(errStr, "uq_customer_nik"):
		res = exception.AlreadyExist("customer")
	case strings.Contains(errStr, "uq_creditlimit_customer_tenor"):
		res = exception.AlreadyExist("credit limits")
	case strings.Contains(errStr, "uq_user_customer"):
		res = exception.AlreadyExist("user customer")
	default:
		res = err
	}

	return
}

func Transaction(ctx context.Context, fn func(context.Context) error) error {
	trx := app.KrediApp.DBConn.Begin()

	ctx = context.WithValue(ctx, constanta.CtxDBTransaction, trx)
	if err := fn(ctx); err != nil {
		trx.Rollback()
		return err
	}

	return trx.Commit().Error
}

func mergedOrNot[T comparable](existing, requested T) T {
	if existing != requested {
		return requested
	} else {
		return existing
	}
}
