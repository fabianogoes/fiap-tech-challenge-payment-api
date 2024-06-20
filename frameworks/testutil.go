package frameworks

import "github.com/gin-gonic/gin"

func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return r
}
