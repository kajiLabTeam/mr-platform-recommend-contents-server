package common

type AbsoluteAddress struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type RecommendContentsRequest struct {
	UserId          string          `json:"userId"`
	AbsoluteAddress AbsoluteAddress `json:"absoluteAddress"`
}

type RecommendContentsResponse struct {
	ContentIds []string `json:"contentIds"`
}
