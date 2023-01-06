package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MigrateSubNetworkInterfaceRequest struct {
	Body *MigrateSubNetworkInterfaceRequestBody `json:"body,omitempty"`
}

func (o MigrateSubNetworkInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateSubNetworkInterfaceRequest struct{}"
	}

	return strings.Join([]string{"MigrateSubNetworkInterfaceRequest", string(data)}, " ")
}
