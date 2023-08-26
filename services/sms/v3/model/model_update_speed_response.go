package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateSpeedResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSpeedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSpeedResponse struct{}"
	}

	return strings.Join([]string{"UpdateSpeedResponse", string(data)}, " ")
}
