package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
