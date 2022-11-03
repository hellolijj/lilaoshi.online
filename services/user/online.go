package user

import (
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
)

// todo cookie 加密存储
func Online(ctx *context.Context) (bool, string) {
	uid := ctx.Input.Session("uid")
	if uid != nil {
		return true, fmt.Sprintf("%v", uid)
	}
	return false, ""
}