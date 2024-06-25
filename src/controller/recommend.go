package controller

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-recommend-contents-server/common"
)

func RecommendContents(c *gin.Context) {
	// リクエストボディを取得
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	body := buf.String()

	// リクエストボディをJSONに変換
	var requestBody map[string]interface{}
	if err := json.Unmarshal([]byte(body), &requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// リクエストボディの中身を確認
	if _, ok := requestBody["userId"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	if _, ok := requestBody["x"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "absoluteAddress.x is required"})
		return
	}

	if _, ok := requestBody["y"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "absoluteAddress.y is required"})
		return
	}

	if _, ok := requestBody["z"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "absoluteAddress.z is required"})
		return
	}

	response := common.RecommendContentsResponse{
		ContentIds: []string{"01F8VYXK67BGC1F9RP1E4S9YK1", "01F8VYXK67BGC1F9RP1E4S9YK2", "01F8VYXK67BGC1F9RP1E4S9YK3"},
	}
	c.JSON(201, gin.H{
		"contentIds": response.ContentIds,
	})
}
