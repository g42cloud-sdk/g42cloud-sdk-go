package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
