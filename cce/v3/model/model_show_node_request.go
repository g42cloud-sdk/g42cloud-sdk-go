package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowNodeRequest struct {
	ClusterId string `json:"cluster_id"`

	NodeId string `json:"node_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`
}

func (o ShowNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeRequest struct{}"
	}

	return strings.Join([]string{"ShowNodeRequest", string(data)}, " ")
}
