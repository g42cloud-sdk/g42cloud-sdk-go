package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateServerMetadataRequest struct {
	ServerId string `json:"server_id"`

	Body *UpdateServerMetadataRequestBody `json:"body,omitempty"`
}

func (o UpdateServerMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerMetadataRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerMetadataRequest", string(data)}, " ")
}
