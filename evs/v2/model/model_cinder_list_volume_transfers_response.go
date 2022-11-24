package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CinderListVolumeTransfersResponse struct {
	Transfers      *[]VolumeTransferSummary `json:"transfers,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CinderListVolumeTransfersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListVolumeTransfersResponse struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTransfersResponse", string(data)}, " ")
}
