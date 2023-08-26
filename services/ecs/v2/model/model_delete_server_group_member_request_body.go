package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
