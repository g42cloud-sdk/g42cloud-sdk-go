package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveNodeResponse struct {
	ApiVersion *string `json:"apiVersion,omitempty"`

	Kind *string `json:"kind,omitempty"`

	Spec *RemoveNodesSpec `json:"spec,omitempty"`

	Status         *TaskStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o RemoveNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveNodeResponse struct{}"
	}

	return strings.Join([]string{"RemoveNodeResponse", string(data)}, " ")
}
