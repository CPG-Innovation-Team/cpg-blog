package middleware

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/internal/auth"
	"cpg-blog/middleware/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// NoCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func NoCache(ctx *gin.Context) {
	ctx.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	ctx.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	ctx.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	ctx.Next()
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options(ctx *gin.Context) {
	if ctx.Request.Method != "OPTIONS" {
		ctx.Next()
	} else {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept, token")
		ctx.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatus(200)
	}
}

// Secure is a middleware function that appends security
// and resource access headers.
func Secure(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("X-Frame-Options", "DENY")
	ctx.Header("X-Content-Type-Options", "nosniff")
	ctx.Header("X-XSS-Protection", "1; mode=block")
	if ctx.Request.TLS != nil {
		ctx.Header("Strict-Transport-Security", "max-age=31536000")
	}
	ctx.Next()

	// Also consider adding Content-Security-Policy headers
	// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
}

// JwtAuth 校验token
func JwtAuth(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token == "" {
		ctx.Abort()
		common.SendResponse(ctx, common.ErrToken, "")
		return
	}
	j := jwt.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		//过期处理
		if err == common.ErrTokenExpired {
			ctx.Abort()
			common.SendResponse(ctx, common.ErrTokenExpired, "")
			return
		}
		//其他错误
		common.SendResponse(ctx, err, "")
		return
	}
	ctx.Set("claims", claims)
}

// PermissionAuth 校验权限
func PermissionAuth(ctx *gin.Context) {
	token, _ := jwt.NewJWT().ParseToken(ctx.Request.Header.Get("token"))
	uid := token.Uid
	path := ctx.Request.RequestURI
	e, _ := auth.GetE(ctx)

	// root用户直接进入
	if token.Root == 1 {
		ctx.Next()
		return
	}
	//TODO 先从redis查询是否存在对应的权限记录

	result, err := e.Enforce(cpgConst.UserPrefix+uid, path, cpgConst.Operate)

	//TODO 如果result == false，需要将user的权限查询记录存入redis

	if !result || err != nil {
		common.SendResponse(ctx, common.ErrAccessDenied, err)
		ctx.Abort()
		return
	}
}
