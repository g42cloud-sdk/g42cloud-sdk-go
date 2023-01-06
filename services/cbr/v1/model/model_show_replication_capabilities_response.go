package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowReplicationCapabilitiesResponse struct {
	Regions        *[]ProtectableReplicationCapabilitiesRespRegion `json:"regions,omitempty"`
	HttpStatusCode int                                             `json:"-"`
}

func (o ShowReplicationCapabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationCapabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ShowReplicationCapabilitiesResponse", string(data)}, " ")
}
