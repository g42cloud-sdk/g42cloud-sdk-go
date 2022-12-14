package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateSubNetworkInterfaceRequest struct {
	Body *CreateSubNetworkInterfaceRequestBody `json:"body,omitempty"`
}

func (o CreateSubNetworkInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubNetworkInterfaceRequest struct{}"
	}

	return strings.Join([]string{"CreateSubNetworkInterfaceRequest", string(data)}, " ")
}
