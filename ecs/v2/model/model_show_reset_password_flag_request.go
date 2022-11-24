package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowResetPasswordFlagRequest struct {
	ServerId string `json:"server_id"`
}

func (o ShowResetPasswordFlagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResetPasswordFlagRequest struct{}"
	}

	return strings.Join([]string{"ShowResetPasswordFlagRequest", string(data)}, " ")
}
