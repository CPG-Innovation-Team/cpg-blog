package jwt

import (
	"cpg-blog/global/common"
	"cpg-blog/internal/oauth"
	"github.com/gin-gonic/gin"
)

// JwtAuth 校验token
func JwtAuth(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.Abort()
		common.SendResponse(c, common.ErrToken, "")
		return
	}
	j := oauth.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		//过期处理
		if err == common.ErrTokenExpired {
			c.Abort()
			common.SendResponse(c, common.ErrTokenExpired, "")
			return
		}
		//其他错误
		common.SendResponse(c, err, "")
		return
	}
	c.Set("claims", claims)
}
