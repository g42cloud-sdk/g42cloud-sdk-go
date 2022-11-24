package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteServerGroupMemberRequestBody struct {
	RemoveMember *ServerGroupMember `json:"remove_member"`
}

func (o DeleteServerGroupMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupMemberRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupMemberRequestBody", string(data)}, " ")
}
