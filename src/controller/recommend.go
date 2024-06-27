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
	var requestBody common.RecommendContentsRequest
	if err := json.Unmarshal([]byte(body), &requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := common.RecommendContentsResponse{
		ContentIds: []string{"22b86d69-9340-4abf-a143-3976877035bf"},
	}
	c.JSON(http.StatusOK, response)
}
