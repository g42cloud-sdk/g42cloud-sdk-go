package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowApiVersionRequest struct {
	Version string `json:"version"`
}

func (o ShowApiVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowApiVersionRequest", string(data)}, " ")
}
