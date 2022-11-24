package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveNodeRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *RemoveNodesTask `json:"body,omitempty"`
}

func (o RemoveNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveNodeRequest struct{}"
	}

	return strings.Join([]string{"RemoveNodeRequest", string(data)}, " ")
}
