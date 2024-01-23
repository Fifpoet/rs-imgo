package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rs-imgo/infra"
	"rs-imgo/util"
)

func UpdateCache(c *gin.Context) {
	user := c.Query("user")
	z := util.Str2Int(c.Param("z"))
	x := util.Str2Int(c.Param("x"))
	y := util.Str2Int(c.Param("y"))
	if user == "" {
		c.JSON(http.StatusOK, "no user code")
		c.Abort()
		return
	}
	//添加到context
	infra.StoreUserState(util.Str2Int(user), infra.UserState{
		Height:  8,
		Width:   8,
		CenterX: x,
		CenterY: y,
		Scale:   z,
		Meta:    nil,
	})
	c.JSON(http.StatusOK, "OK")
}
