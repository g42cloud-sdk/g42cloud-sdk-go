package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNodeResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *NodeMetadata `json:"metadata,omitempty"`

	Spec *NodeSpec `json:"spec,omitempty"`

	Status         *NodeStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeResponse struct{}"
	}

	return strings.Join([]string{"UpdateNodeResponse", string(data)}, " ")
}
