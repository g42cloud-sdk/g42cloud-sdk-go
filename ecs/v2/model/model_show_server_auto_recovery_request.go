package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowServerAutoRecoveryRequest struct {
	ServerId string `json:"server_id"`
}

func (o ShowServerAutoRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerAutoRecoveryRequest struct{}"
	}

	return strings.Join([]string{"ShowServerAutoRecoveryRequest", string(data)}, " ")
}
