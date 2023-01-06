package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PolicyUpdate struct {
	Enabled *bool `json:"enabled,omitempty"`

	Name *string `json:"name,omitempty"`

	OperationDefinition *PolicyoOdCreate `json:"operation_definition,omitempty"`

	Trigger *PolicyTriggerReq `json:"trigger,omitempty"`
}

func (o PolicyUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyUpdate struct{}"
	}

	return strings.Join([]string{"PolicyUpdate", string(data)}, " ")
}
