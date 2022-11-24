package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CinderShowVolumeTransferRequest struct {
	TransferId string `json:"transfer_id"`
}

func (o CinderShowVolumeTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderShowVolumeTransferRequest struct{}"
	}

	return strings.Join([]string{"CinderShowVolumeTransferRequest", string(data)}, " ")
}
