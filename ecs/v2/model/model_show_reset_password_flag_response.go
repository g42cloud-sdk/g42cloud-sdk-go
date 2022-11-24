package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowResetPasswordFlagResponse struct {
	ResetpwdFlag   *string `json:"resetpwd_flag,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResetPasswordFlagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResetPasswordFlagResponse struct{}"
	}

	return strings.Join([]string{"ShowResetPasswordFlagResponse", string(data)}, " ")
}
