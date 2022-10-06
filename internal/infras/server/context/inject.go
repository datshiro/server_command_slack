package context

import (
	"github.com/gin-gonic/gin"
)

const (
	injectKey = "inject_context"
)

func (c Context) GinInject(g *gin.Context) {
	g.Set(injectKey, c)
}

func GetContext(g *gin.Context) Context {
	val, ok := g.Get(injectKey)
	if !ok {
		panic("empty context")
	}
	return val.(Context)
}

// func GetRedis(g *gin.Context) redis.Redis {
// 	appCtx := GetContext(g)
// 	return appCtx.Redis
// }
