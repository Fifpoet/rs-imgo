package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rs-imgo/infra"
	"rs-imgo/util"
)

func UpdateCache(c *gin.Context) {
	user := c.Query("user-code")
	x := c.Query("x")
	y := c.Query("y")
	if user == "" {
		c.JSON(http.StatusBadRequest, "no user code")
		c.Abort()
	}
	//添加到context
	infra.StoreUserState(util.Str2Int(user), infra.UserState{
		Height:  8,
		Width:   8,
		CenterX: util.Str2Int(x),
		CenterY: util.Str2Int(y),
		Meta:    nil,
	})
	c.JSON(http.StatusOK, "OK")
}
