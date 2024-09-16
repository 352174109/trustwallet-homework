package utils

import (
	"context"
	"fmt"
	"reflect"
	"runtime"

	"github.com/352174109/trustwallet-homework/internal/logs"
)

func WrapRecover(ctx context.Context, process func() error) {
	defer func() {
		if err := recover(); err != nil {
			msg := fmt.Sprintf("Process run %s failed, %s", getFunctionName(process), err)
			logs.CtxError(ctx, msg)
		}
	}()

	if err := process(); err != nil {
		logs.CtxError(ctx, "execute %s err: %+v", getFunctionName(process), err)
	}
}

func getFunctionName(i interface{}) string {
	if i == nil {
		return "nil"
	}

	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
