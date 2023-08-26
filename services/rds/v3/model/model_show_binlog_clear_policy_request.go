package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowBinlogClearPolicyRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o ShowBinlogClearPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBinlogClearPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowBinlogClearPolicyRequest", string(data)}, " ")
}
