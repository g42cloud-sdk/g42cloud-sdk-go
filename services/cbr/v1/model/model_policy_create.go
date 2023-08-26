package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PolicyCreate struct {
	Enabled *bool `json:"enabled,omitempty"`

	Name string `json:"name"`

	OperationDefinition *PolicyoOdCreate `json:"operation_definition"`

	OperationType string `json:"operation_type"`

	Trigger *PolicyTriggerReq `json:"trigger"`
}

func (o PolicyCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyCreate struct{}"
	}

	return strings.Join([]string{"PolicyCreate", string(data)}, " ")
}
