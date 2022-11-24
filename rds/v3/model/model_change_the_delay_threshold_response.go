package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeTheDelayThresholdResponse struct {
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeTheDelayThresholdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeTheDelayThresholdResponse struct{}"
	}

	return strings.Join([]string{"ChangeTheDelayThresholdResponse", string(data)}, " ")
}
