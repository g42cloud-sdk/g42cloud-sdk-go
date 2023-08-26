package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowNodePoolRequest struct {
	ClusterId string `json:"cluster_id"`

	NodepoolId string `json:"nodepool_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`
}

func (o ShowNodePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodePoolRequest struct{}"
	}

	return strings.Join([]string{"ShowNodePoolRequest", string(data)}, " ")
}
