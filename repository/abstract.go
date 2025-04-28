package repository

import (
	"context"
	"gorm.io/gorm"
	"kredi-plus.com/be/app"
	"kredi-plus.com/be/lib/constanta"
)

func GetDB(ctx context.Context) (tx *gorm.DB) {
	tx, ok := ctx.Value(constanta.CtxDBTransaction).(*gorm.DB)
	if !ok {
		tx = app.KrediApp.DBConn
	}
	return
}

func GetDBWithStatus(ctx context.Context) (tx *gorm.DB, isTransaction bool) {
	tx, isTransaction = ctx.Value(constanta.CtxDBTransaction).(*gorm.DB)
	if !isTransaction {
		tx = app.KrediApp.DBConn
	}
	return
}
