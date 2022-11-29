package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecurityPolicyRequestBody struct {
	SecurityPolicy *UpdateSecurityPolicyOption `json:"security_policy"`
}

func (o UpdateSecurityPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolicyRequestBody", string(data)}, " ")
}
