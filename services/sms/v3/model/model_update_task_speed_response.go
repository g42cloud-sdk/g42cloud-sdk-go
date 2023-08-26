package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTaskSpeedResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTaskSpeedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskSpeedResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskSpeedResponse", string(data)}, " ")
}
