package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowOffSiteBackupPolicyResponse struct {
	PolicyPara     *[]GetOffSiteBackupPolicy `json:"policy_para,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowOffSiteBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOffSiteBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowOffSiteBackupPolicyResponse", string(data)}, " ")
}
