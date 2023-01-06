package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateMemberStatusRequest struct {
	MemberId string `json:"member_id"`

	BackupId string `json:"backup_id"`

	Body *UpdateMember `json:"body,omitempty"`
}

func (o UpdateMemberStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateMemberStatusRequest", string(data)}, " ")
}
