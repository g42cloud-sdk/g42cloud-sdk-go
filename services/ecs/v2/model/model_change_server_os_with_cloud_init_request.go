package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ChangeServerOsWithCloudInitRequest struct {
	ServerId string `json:"server_id"`

	Body *ChangeServerOsWithCloudInitRequestBody `json:"body,omitempty"`
}

func (o ChangeServerOsWithCloudInitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithCloudInitRequest struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithCloudInitRequest", string(data)}, " ")
}
