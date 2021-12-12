package middlewares

import (
	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	UserId   int      `json:"user_id"`
	UserCode string   `json:"user_code"`
	UserName string   `json:"user_name"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	IsSuper  bool     `json:"is_super"`
	IsOps    bool     `json:"is_ops"`
	IsLeader bool     `json:"is_leader"`
}

var (
	loginUser *LoginUser
)

// 判断用户书否登陆中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginUser = &LoginUser{
			UserId: 1,
			UserCode: "jian.li",
			UserName: "李健",
			Phone: "xxx",
			Email: "xxx",
			IsSuper: true,
			IsOps: true,
			IsLeader: true,
		}
	}
}

func GetLoginUser() *LoginUser {
	return loginUser
}