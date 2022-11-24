package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeCreateRequest struct {
	Kind string `json:"kind"`

	ApiVersion string `json:"apiVersion"`

	Metadata *NodeMetadata `json:"metadata,omitempty"`

	Spec *NodeSpec `json:"spec"`
}

func (o NodeCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeCreateRequest struct{}"
	}

	return strings.Join([]string{"NodeCreateRequest", string(data)}, " ")
}
