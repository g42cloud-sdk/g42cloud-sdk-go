package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailoverModeRequest struct {
	Mode string `json:"mode"`
}

func (o FailoverModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailoverModeRequest struct{}"
	}

	return strings.Join([]string{"FailoverModeRequest", string(data)}, " ")
}
