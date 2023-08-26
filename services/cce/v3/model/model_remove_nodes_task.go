package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RemoveNodesTask struct {
	ApiVersion *string `json:"apiVersion,omitempty"`

	Kind *string `json:"kind,omitempty"`

	Spec *RemoveNodesSpec `json:"spec"`

	Status *TaskStatus `json:"status,omitempty"`
}

func (o RemoveNodesTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveNodesTask struct{}"
	}

	return strings.Join([]string{"RemoveNodesTask", string(data)}, " ")
}
