package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
