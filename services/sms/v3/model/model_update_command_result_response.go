package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateCommandResultResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCommandResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCommandResultResponse struct{}"
	}

	return strings.Join([]string{"UpdateCommandResultResponse", string(data)}, " ")
}
