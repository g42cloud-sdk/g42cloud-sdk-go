package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NovaListServerSecurityGroupsRequest struct {
	ServerId string `json:"server_id"`
}

func (o NovaListServerSecurityGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListServerSecurityGroupsRequest struct{}"
	}

	return strings.Join([]string{"NovaListServerSecurityGroupsRequest", string(data)}, " ")
}
