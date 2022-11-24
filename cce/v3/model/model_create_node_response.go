package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNodeResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *NodeMetadata `json:"metadata,omitempty"`

	Spec *NodeSpec `json:"spec,omitempty"`

	Status         *NodeStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodeResponse struct{}"
	}

	return strings.Join([]string{"CreateNodeResponse", string(data)}, " ")
}
