package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListNodePoolsRequest struct {
	ClusterId string `json:"cluster_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`

	ShowDefaultNodePool *string `json:"showDefaultNodePool,omitempty"`
}

func (o ListNodePoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodePoolsRequest struct{}"
	}

	return strings.Join([]string{"ListNodePoolsRequest", string(data)}, " ")
}
