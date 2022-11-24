package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CinderAcceptVolumeTransferRequestBody struct {
	Accept *CinderAcceptVolumeTransferOption `json:"accept"`
}

func (o CinderAcceptVolumeTransferRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderAcceptVolumeTransferRequestBody struct{}"
	}

	return strings.Join([]string{"CinderAcceptVolumeTransferRequestBody", string(data)}, " ")
}
