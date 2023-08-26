package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SetOffSiteBackupPolicyRequestBody struct {
	PolicyPara []OffSiteBackupPolicy `json:"policy_para"`
}

func (o SetOffSiteBackupPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetOffSiteBackupPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetOffSiteBackupPolicyRequestBody", string(data)}, " ")
}
