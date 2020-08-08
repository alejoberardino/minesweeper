package routes

import (
	"github.com/alejoberardino/minesweeper/model"
	"github.com/gin-gonic/gin"
)

// test godoc
// @Summary Test endpoint
// @Description return sample message
// @Produce  json
// @Success 200 {object} model.MessageResponse
// @Router /accounts [get]
func Test(c *gin.Context) {
	c.JSON(200, model.MessageResponse{
		Message: "Hello there!",
	})
}
