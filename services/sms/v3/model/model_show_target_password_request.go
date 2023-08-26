package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowTargetPasswordRequest struct {
	Id string `json:"id"`
}

func (o ShowTargetPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTargetPasswordRequest struct{}"
	}

	return strings.Join([]string{"ShowTargetPasswordRequest", string(data)}, " ")
}
