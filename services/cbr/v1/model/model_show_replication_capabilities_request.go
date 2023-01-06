package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowReplicationCapabilitiesRequest struct {
}

func (o ShowReplicationCapabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationCapabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ShowReplicationCapabilitiesRequest", string(data)}, " ")
}
