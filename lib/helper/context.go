package helper

import (
	"context"
	"kredi-plus.com/be/dto/app_model"
	"kredi-plus.com/be/lib/constanta"
)

func GetAuthSessionModel(ctx context.Context) app_model.UserSession {
	model, _ := ctx.Value(constanta.CtxNameAuthAccess).(app_model.UserSession)
	return model
}

func SetAuthSessionModel(ctx context.Context, model app_model.UserSession) context.Context {
	return context.WithValue(ctx, constanta.CtxNameAuthAccess, model)
}
