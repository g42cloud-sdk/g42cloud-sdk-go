package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteAnimatedGraphicsTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAnimatedGraphicsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAnimatedGraphicsTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteAnimatedGraphicsTaskResponse", string(data)}, " ")
}
