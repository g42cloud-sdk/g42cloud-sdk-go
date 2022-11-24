package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddServerGroupMemberRequest struct {
	ServerGroupId string `json:"server_group_id"`

	Body *AddServerGroupMemberRequestBody `json:"body,omitempty"`
}

func (o AddServerGroupMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServerGroupMemberRequest struct{}"
	}

	return strings.Join([]string{"AddServerGroupMemberRequest", string(data)}, " ")
}
