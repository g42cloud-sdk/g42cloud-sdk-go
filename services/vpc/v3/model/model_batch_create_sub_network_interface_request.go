package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateSubNetworkInterfaceRequest struct {
	Body *BatchCreateSubNetworkInterfaceRequestBody `json:"body,omitempty"`
}

func (o BatchCreateSubNetworkInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubNetworkInterfaceRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSubNetworkInterfaceRequest", string(data)}, " ")
}
