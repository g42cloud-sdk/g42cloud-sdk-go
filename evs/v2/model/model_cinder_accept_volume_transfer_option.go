package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CinderAcceptVolumeTransferOption struct {
	AuthKey string `json:"auth_key"`
}

func (o CinderAcceptVolumeTransferOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderAcceptVolumeTransferOption struct{}"
	}

	return strings.Join([]string{"CinderAcceptVolumeTransferOption", string(data)}, " ")
}
