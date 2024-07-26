package handler

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"option_get/log"
)

func requestSuccessful(ctx iris.Context, code int64, msg string, data any) {
	ctx.StatusCode(http.StatusOK)
	if err := ctx.JSON(map[string]any{
		"code": code,
		"msg":  msg,
		"data": data,
	}); err != nil {
		log.Debug(err.Error())
	}
}

func requsetFailed(ctx iris.Context, err error) {
	ctx.StatusCode(http.StatusInternalServerError)
	if err := ctx.JSON(map[string]any{
		"code": -1,
		"msg":  err.Error(),
		"data": nil,
	}); err != nil {
		log.Debug(err.Error())
	}
}
