package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AllowDbUserPrivilegeRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *GrantRequest `json:"body,omitempty"`
}

func (o AllowDbUserPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowDbUserPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"AllowDbUserPrivilegeRequest", string(data)}, " ")
}
