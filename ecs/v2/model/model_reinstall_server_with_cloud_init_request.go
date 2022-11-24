package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReinstallServerWithCloudInitRequest struct {
	ServerId string `json:"server_id"`

	Body *ReinstallServerWithCloudInitRequestBody `json:"body,omitempty"`
}

func (o ReinstallServerWithCloudInitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerWithCloudInitRequest struct{}"
	}

	return strings.Join([]string{"ReinstallServerWithCloudInitRequest", string(data)}, " ")
}
