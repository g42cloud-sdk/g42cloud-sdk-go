package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddMemberRequest struct {
	BackupId string `json:"backup_id"`

	Body *AddMembersReq `json:"body,omitempty"`
}

func (o AddMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMemberRequest struct{}"
	}

	return strings.Join([]string{"AddMemberRequest", string(data)}, " ")
}
