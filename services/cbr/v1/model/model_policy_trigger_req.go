package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PolicyTriggerReq struct {
	Properties *PolicyTriggerPropertiesReq `json:"properties"`
}

func (o PolicyTriggerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTriggerReq struct{}"
	}

	return strings.Join([]string{"PolicyTriggerReq", string(data)}, " ")
}
