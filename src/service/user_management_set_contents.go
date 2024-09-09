package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/kajiLabTeam/mr-platform-recommend-contents-server/common"
)

func UserManagementSetContents(userId string, contentIds []string) (common.ResponseUserManagementSetContents, error) {
	userManagementServerUrl := os.Getenv("USER_SERVER_URL")
	endPoint := userManagementServerUrl + "/api/content/set"
	requestUserManagementSetContents := common.RequestUserManagementSetContents{
		UserId:     userId,
		ContentIds: contentIds,
	}
	requestUserManagementSetContentsStr, err := json.Marshal(requestUserManagementSetContents)
	if err != nil {
		return common.ResponseUserManagementSetContents{}, err
	}

	// リクエストを作成
	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(requestUserManagementSetContentsStr))
	if err != nil {
		return common.ResponseUserManagementSetContents{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return common.ResponseUserManagementSetContents{}, err
	}
	defer response.Body.Close()

	// 201以外はエラー
	if response.StatusCode != http.StatusCreated {
		var responseError common.ResponseError
		err = json.NewDecoder(response.Body).Decode(&responseError)
		if err != nil {
			return common.ResponseUserManagementSetContents{}, err
		}
		return common.ResponseUserManagementSetContents{}, errors.New(responseError.ErrorMessage)
	}

	// レスポンスをパース
	var responseUserManagementSetContents common.ResponseUserManagementSetContents
	err = json.NewDecoder(response.Body).Decode(&responseUserManagementSetContents)
	if err != nil {
		return common.ResponseUserManagementSetContents{}, err
	}

	return responseUserManagementSetContents, nil
}
