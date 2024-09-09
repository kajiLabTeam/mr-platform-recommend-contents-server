package common

import "github.com/uber/h3-go/v4"

type UserLocation struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Height float64 `json:"height"`
}

type RequestRecomendContent struct {
	UserId       string       `json:"userId"`
	UserLocation UserLocation `json:"userLocation"`
}

type RequestUserManagementSetContents struct {
	UserId     string   `json:"userId"`
	ContentIds []string `json:"contentIds"`
}

type ResponseError struct {
	ErrorMessage string `json:"error"`
}

type ResponseUserManagementSetContents struct {
	ContentIds []string `json:"contentIds"`
}

type CellDistance struct {
	Cell     h3.Cell
	Distance float64
}
