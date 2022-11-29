package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowMemberDetailRequest struct {
	BackupId string `json:"backup_id"`

	MemberId string `json:"member_id"`
}

func (o ShowMemberDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberDetailRequest", string(data)}, " ")
}
