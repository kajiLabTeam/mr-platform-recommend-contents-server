package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-recommend-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-recommend-contents-server/model"
	"github.com/kajiLabTeam/mr-platform-recommend-contents-server/service"
)

func RecomendContents(c *gin.Context) {
	var req common.RequestRecomendContent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	predictUserNextCells := service.PredictUserNextCells(req.UserLocation.Lat, req.UserLocation.Lon, 10)

	// predictUserNextCells に含まれるコンテンツを取得する
	contentIds, err := model.H3CellsToContentIds(predictUserNextCells)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ユーザーサーバーにコンテンツを返す
	reaContentIds, err := service.UserManagementSetContents(req.UserId, contentIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reaContentIds)
}
